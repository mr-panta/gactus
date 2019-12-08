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
	addrToConnConfigMap map[string]*pb.ConnectionConfig
	healthCheckInterval time.Duration
	lock                sync.RWMutex
}

func newServiceManager(healthCheckInterval int) *serviceManager {
	m := &serviceManager{
		routeToCommandMap:   make(map[string]string),
		commandToAddrsMap:   make(map[string][]string),
		addrToClientMap:     make(map[string]tcpclient.Client),
		addrToConnConfigMap: make(map[string]*pb.ConnectionConfig),
		healthCheckInterval: time.Duration(healthCheckInterval) * time.Second,
	}
	go m.startServiceDoctor(true)
	return m
}

// Getter and setter methods

func (m *serviceManager) getCommandByRoute(route string) (command string, exists bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	command, exists = m.routeToCommandMap[route]
	return command, exists
}

func (m *serviceManager) getCommandsWithAddrsList() (commands []string, addrsList [][]string) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for command, addrs := range m.commandToAddrsMap {
		commands = append(commands, command)
		addrsList = append(addrsList, addrs)
	}
	return commands, addrsList
}

func (m *serviceManager) getAddrsByCommand(command string) (addrs []string, exists bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	addrs, exists = m.commandToAddrsMap[command]
	return addrs, exists
}
func (m *serviceManager) getAddrsWithClients() (addrs []string, clients []tcpclient.Client) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for addr, client := range m.addrToClientMap {
		addrs = append(addrs, addr)
		clients = append(clients, client)
	}
	return addrs, clients
}

func (m *serviceManager) getClientByAddr(addr string) (client tcpclient.Client, exists bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	client, exists = m.addrToClientMap[addr]
	return client, exists
}

func (m *serviceManager) getAddrsWithConnConfigs() (addrs []string, cfgs []*pb.ConnectionConfig) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for addr, cfg := range m.addrToConnConfigMap {
		addrs = append(addrs, addr)
		cfgs = append(cfgs, cfg)
	}
	return addrs, cfgs
}

// Normal methods

func (m *serviceManager) getCommand(method, path string) (command string, exists bool) {
	route := getRoute(method, path)
	return m.getCommandByRoute(route)
}

func (m *serviceManager) getServiceConn(command string) (service tcpclient.Client, exists bool) {
	addrs, _ := m.getAddrsByCommand(command)
	addrsLength := len(addrs)
	if addrsLength == 0 {
		return nil, false
	}
	addr := addrs[rand.Intn(addrsLength)]
	return m.getClientByAddr(addr)
}

func (m *serviceManager) startServiceDoctor(loop bool) {
	ctx := logger.GetContextWithLogID(context.Background(), "service_doctor")
	for {
		if loop {
			time.Sleep(m.healthCheckInterval)
		}
		addrs, clients := m.getAddrsWithClients()
		for i, addr := range addrs {
			client := clients[i]
			_, err := serviceHealthCheck(ctx, client)
			if err != nil {
				m.abandonService(addr)
			}
		}
		if !loop {
			return
		}
	}
}

func serviceHealthCheck(ctx context.Context, client tcpclient.Client) (res *pb.HealthCheckResponse, err error) {
	req := &pb.HealthCheckRequest{
		Address: client.GetHostAddr(),
	}
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
		logger.Errorf(ctx, err.Error())
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
	ctx := logger.GetContextWithLogID(context.Background(), "abandon_service")
	logger.Debugf(ctx, "abandon service[%s]", addr)
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
	commands, addrsList := m.getCommandsWithAddrsList()
	for i, command := range commands {
		addrs := addrsList[i]
		for _, addr := range addrs {
			req.CommandAddressPairs = append(req.CommandAddressPairs, &pb.CommandAddressPair{
				Command: command,
				Address: addr,
			})
		}
	}
	addrs, cfgs := m.getAddrsWithConnConfigs()
	for i, addr := range addrs {
		cfg := cfgs[i]
		req.AddrConfigs = append(req.AddrConfigs, &pb.AddressConfig{
			Address:    addr,
			ConnConfig: cfg,
		})
	}
	data, err := wrapAndMarshalRequest(ctx, config.CMDServiceUpdateRegistries, req)
	if err != nil {
		return err
	}
	_, clients := m.getAddrsWithClients()
	for _, client := range clients {
		err = updateProcessorRegistries(client, data)
		if err != nil {
			logger.Errorf(ctx, err.Error())
		}
	}
	return nil
}
