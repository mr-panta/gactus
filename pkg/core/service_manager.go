package core

import (
	"fmt"
	"math/rand"
	"strings"

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

func (m *serviceManager) registerProcessors(req *pb.RegisterProcessorsRequest, res *pb.RegisterProcessorsResponse) (code uint32) {
	for _, registry := range req.ProcessorRegistries {
		if registry.HttpConfig != nil {
			method := getMethodString(registry.HttpConfig.Method)
			route := getRoute(method, registry.HttpConfig.Path)
			m.routeToCommandMap[route] = registry.Command
		}
		m.commandToAddrsMap[registry.Command] = append(m.commandToAddrsMap[registry.Command], req.Addr)
	}
	return uint32(pb.Constant_RESPONSE_OK)
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
