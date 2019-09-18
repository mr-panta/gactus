package core

import (
	"fmt"
	"math/rand"

	"github.com/mr-panta/gactus/pkg/tcp"
)

type serviceManager struct {
	routeToCommandMap map[string]string // key format = `method:path`
	commandToAddrsMap map[string][]string
	addrToConnMap     map[string]tcp.Client
}

func newServiceManager() *serviceManager {
	return &serviceManager{
		routeToCommandMap: make(map[string]string),
		commandToAddrsMap: make(map[string][]string),
	}
}

func (m *serviceManager) getCommand(method, path string) (command string, exists bool) {
	route := fmt.Sprintf("%s:%s", method, path)
	command, exists = m.routeToCommandMap[route]
	return
}

func (m *serviceManager) getServiceConn(command string) (service tcp.Client, exists bool) {
	addrs := m.commandToAddrsMap[command]
	addrsLength := len(addrs)
	if addrsLength == 0 {
		return nil, false
	}
	addr := addrs[rand.Intn(addrsLength)]
	service, exists = m.addrToConnMap[addr]
	return
}
