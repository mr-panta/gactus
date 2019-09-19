package service

import (
	"net"

	"github.com/golang/protobuf/proto"
)

// Handler [TOWRITE]
type Handler interface {
	SetProcess(command string, process func(req, res proto.Message) (code uint32))
	SendCoreRequest(logID, command string, req, res proto.Message) (code uint32, err error)
	ServeTCP(conn net.Conn)
}
