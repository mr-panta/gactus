package core

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/mr-panta/gactus/internal/config"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/go-logger"

	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-tcpclient"
)

type serviceManager struct {
	routeToCommandMap   map[string]string
	commandToAddrsMap   map[string][]string
	addrToClientMap     map[string]tcpclient.Client
	addrToActiveTimeMap map[string]time.Time
	addrToConnConfigMap map[string]*pb.ConnectionConfig
	healthCheckInterval time.Duration
	lock                sync.Mutex
}

func newServiceManager(healthCheckInterval int) *serviceManager {
	m := &serviceManager{
		routeToCommandMap:   make(map[string]string),
		commandToAddrsMap:   make(map[string][]string),
		addrToClientMap:     make(map[string]tcpclient.Client),
		addrToActiveTimeMap: make(map[string]time.Time),
		addrToConnConfigMap: make(map[string]*pb.ConnectionConfig),
		healthCheckInterval: time.Duration(healthCheckInterval) * time.Second,
	}
	go m.startServiceDoctor()
	return m
}

func (m *serviceManager) getCommand(method, path string) (command string, exists bool) {
	route := getRoute(method, path)
	command, exists = m.routeToCommandMap[route]
	return
}

func (m *serviceManager) getServiceConn(command string) (service tcpclient.Client, exists bool) {
	addrs := m.commandToAddrsMap[command]
	addrsLength := len(addrs)
	if addrsLength == 0 {
		return nil, false
	}
	addr := addrs[rand.Intn(addrsLength)]
	service, exists = m.addrToClientMap[addr]
	return
}

func (m *serviceManager) startServiceDoctor() {
	ctx := logger.GetContextWithLogID(context.Background(), "service_doctor")
	for {
		time.Sleep(m.healthCheckInterval)
		for addr, client := range m.addrToClientMap {
			_, err := serviceHealthCheck(ctx, client)
			if err != nil {
				m.abandonService(addr)
			}
		}
	}
}

func serviceHealthCheck(ctx context.Context, client tcpclient.Client) (res *pb.HealthCheckResponse, err error) {
	req := &pb.HealthCheckRequest{}
	res = &pb.HealthCheckResponse{}
	body, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	wrappedReq := &pb.Request{
		LogId:   logger.GetLogID(ctx),
		Command: config.CMDServiceHealthCheck,
		IsProto: true,
		Body:    body,
	}
	input, err := proto.Marshal(wrappedReq)
	if err != nil {
		return nil, err
	}
	output, err := client.Send(input)
	if err != nil {
		return nil, err
	}
	code, err := unmarshalAndUnwrapResponse(output, res)
	if err != nil {
		return nil, err
	}
	if code != uint32(pb.Constant_RESPONSE_OK) {
		return nil, errors.New(res.DebugMessage)
	}
	return res, nil
}

func verifyServiceAddresses(ctx context.Context, addrs []string) (addr string, err error) {
	for _, addr := range addrs {
		client, err := tcpclient.NewClient(addr, 1, 1, 0, 0, time.Second)
		if err == nil {
			_, err = serviceHealthCheck(ctx, client)
			if err == nil {
				return addr, nil
			}
		}
		client.Close()
	}
	return "", errors.New("cannot verify service addresses")
}

func (m *serviceManager) registerService(ctx context.Context, wrappedReq *pb.Request) (wrappedRes *pb.Response, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	req := &pb.RegisterServiceRequest{}
	res := &pb.RegisterServiceResponse{}
	err = proto.Unmarshal(wrappedReq.Body, req)
	if err != nil {
		return nil, err
	}
	// Check usable IP Address
	serviceAddr, err := verifyServiceAddresses(ctx, req.Addresses)
	if err != nil {
		return nil, err
	}
	// Setup connection to service
	m.addrToConnConfigMap[serviceAddr] = req.ConnConfig
	client, exists := m.addrToClientMap[serviceAddr]
	if exists {
		client.Close()
	}
	client, err = tcpclient.NewClient(
		serviceAddr,
		int(req.ConnConfig.MinConns),
		int(req.ConnConfig.MaxConns),
		time.Duration(req.ConnConfig.IdleConnTimeout)*time.Millisecond,
		time.Duration(req.ConnConfig.WaitConnTimeout)*time.Millisecond,
		time.Duration(req.ConnConfig.ClearPeriod)*time.Millisecond,
	)
	if err != nil {
		return nil, err
	}
	m.addrToClientMap[serviceAddr] = client
	m.addrToActiveTimeMap[serviceAddr] = time.Now()
	// Update processors registries
	for _, registry := range req.ProcessorRegistries {
		if registry.HttpConfig != nil {
			method := getMethodString(registry.HttpConfig.Method)
			route := getRoute(method, registry.HttpConfig.Path)
			m.routeToCommandMap[route] = registry.Command
		}
		logger.Debugf(ctx, "register command[%s] from address[%s]", registry.Command, serviceAddr)
		// Check existing address and commend before adding it
		isNewAddr := true
		for _, addr := range m.commandToAddrsMap[registry.Command] {
			if addr == serviceAddr {
				isNewAddr = false
				break
			}
		}
		if isNewAddr {
			m.commandToAddrsMap[registry.Command] = append(m.commandToAddrsMap[registry.Command], serviceAddr)
		}
	}
	res.Address = serviceAddr
	wrappedRes = &pb.Response{}
	wrappedRes.Body, err = proto.Marshal(res)
	wrappedRes.Code = uint32(pb.Constant_RESPONSE_OK)
	if err != nil {
		return nil, err
	}
	return wrappedRes, nil
}

func getRoute(method, path string) string {
	return fmt.Sprintf("%s:%s", strings.ToUpper(method), path)
}

func getMethodString(method pb.Constant_HttpMethod) string {
	switch method {
	case pb.Constant_HTTP_METHOD_GET:
		return "GET"
	case pb.Constant_HTTP_METHOD_POST:
		return "POST"
	}
	return "UNKNOWN"
}

func (m *serviceManager) abandonService(addr string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for command, addrs := range m.commandToAddrsMap {
		exists := false
		idx := 0
		a := ""
		for idx, a = range addrs {
			exists = exists || (addr == a)
			if exists {
				break
			}
		}
		if exists {
			if len(m.commandToAddrsMap[command]) > 1 {
				m.commandToAddrsMap[command] = append(
					m.commandToAddrsMap[command][:idx],
					m.commandToAddrsMap[command][idx+1:]...,
				)
			} else {
				delete(m.commandToAddrsMap, command)
				for route, cmd := range m.routeToCommandMap {
					if cmd == command {
						delete(m.routeToCommandMap, route)
					}
				}
			}
		}
	}
	if client, exists := m.addrToClientMap[addr]; exists {
		client.Close()
		delete(m.addrToClientMap, addr)
	}
	if _, exists := m.addrToActiveTimeMap[addr]; exists {
		delete(m.addrToActiveTimeMap, addr)
	}
}

func wrapAndMarshalRequest(ctx context.Context, command string, req proto.Message) (data []byte, err error) {
	body, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	logID := logger.GetLogID(ctx)
	wrappedReq := &pb.Request{
		Command: command,
		LogId:   logID,
		Body:    body,
		IsProto: true,
	}
	data, err = proto.Marshal(wrappedReq)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func unmarshalAndUnwrapResponse(data []byte, res proto.Message) (code uint32, err error) {
	wrappedRes := &pb.Response{}
	err = proto.Unmarshal(data, wrappedRes)
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_UNPACK_RESPONSE), err
	}
	err = proto.Unmarshal(wrappedRes.Body, res)
	if err != nil {
		return wrappedRes.Code, err
	}
	return wrappedRes.Code, nil
}

func updateProcessorRegistries(client tcpclient.Client, data []byte) (err error) {
	res := &pb.UpdateRegistriesResponse{}
	output, err := client.Send(data)
	if err != nil {
		return fmt.Errorf("cannot send processor registries to service[%s]: %v", client.GetHostAddr(), err)
	}
	code, err := unmarshalAndUnwrapResponse(output, res)
	if err != nil {
		return fmt.Errorf("cannot unpack wrapped response: %v", err)
	}
	if code != uint32(pb.Constant_RESPONSE_OK) {
		return fmt.Errorf("code[%d]: %s", code, res.DebugMessage)
	}
	return nil
}

func (m *serviceManager) broadcastProcessorRegistries(ctx context.Context) (err error) {
	logger.Debugf(ctx, "broadcast processor registries")
	req := &pb.UpdateRegistriesRequest{}
	for command, addrs := range m.commandToAddrsMap {
		for _, addr := range addrs {
			req.CommandAddressPairs = append(req.CommandAddressPairs, &pb.CommandAddressPair{
				Command: command,
				Address: addr,
			})
		}
	}
	for addr, connConfig := range m.addrToConnConfigMap {
		req.AddrConfigs = append(req.AddrConfigs, &pb.AddressConfig{
			Address:    addr,
			ConnConfig: connConfig,
		})
	}
	data, err := wrapAndMarshalRequest(ctx, config.CMDServiceUpdateRegistries, req)
	if err != nil {
		return err
	}
	for _, client := range m.addrToClientMap {
		err = updateProcessorRegistries(client, data)
		if err != nil {
			logger.Errorf(ctx, err.Error())
		}
	}
	return nil
}
