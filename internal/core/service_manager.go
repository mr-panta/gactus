package core

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/go-logger"

	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-tcpclient"
)

type serviceManager struct {
	routeToCommandMap map[string]string // key format = `method:path`
	commandToAddrsMap map[string][]string
	addrToConnMap     map[string]tcpclient.Client
}

func newServiceManager() *serviceManager {
	return &serviceManager{
		routeToCommandMap: make(map[string]string),
		commandToAddrsMap: make(map[string][]string),
		addrToConnMap:     make(map[string]tcpclient.Client),
	}
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
	service, exists = m.addrToConnMap[addr]
	return
}

func (m *serviceManager) registerProcessors(ctx context.Context, wrappedReq *pb.Request) (wrappedRes *pb.Response, err error) {
	req := &pb.RegisterProcessorsRequest{}
	res := &pb.RegisterProcessorsResponse{}
	err = proto.Unmarshal(wrappedReq.Body, req)
	if err != nil {
		return nil, err
	}
	for _, registry := range req.ProcessorRegistries {
		if registry.HttpConfig != nil {
			method := getMethodString(registry.HttpConfig.Method)
			route := getRoute(method, registry.HttpConfig.Path)
			m.routeToCommandMap[route] = registry.Command
		}
		logger.Debugf(ctx, "register command[%s] from address[%s]", registry.Command, req.Addr)
		// Check existing address and commend before adding it
		isNewAddr := true
		for _, addr := range m.commandToAddrsMap[registry.Command] {
			if addr == req.Addr {
				isNewAddr = false
				break
			}
		}
		if isNewAddr {
			m.commandToAddrsMap[registry.Command] = append(m.commandToAddrsMap[registry.Command], req.Addr)
		}
	}
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