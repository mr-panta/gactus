package service

import (
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/go-tcpclient"
)

// Handler [TOWRITE]
type Handler interface {
	SetProcess(command string, process func(req, res proto.Message) (code uint32))
	SendCoreRequest(logID, command string, req, res proto.Message) (code uint32, err error)
	ServeTCP(conn net.Conn)
}

// NewHandler [TOWRITE]
func NewHandler(coreAddr string, minConns, maxConns, idleConnTimeout, waitConnTimeout, clearPeriod int) (Handler, error) {
	coreClient, err := tcpclient.NewClient(
		coreAddr,
		minConns,
		maxConns,
		time.Duration(idleConnTimeout)*time.Millisecond,
		time.Duration(waitConnTimeout)*time.Millisecond,
		time.Duration(clearPeriod)*time.Millisecond,
	)
	if err != nil {
		return nil, err
	}
	return &handler{
		coreClient:        coreClient,
		commandProcessMap: make(map[string]func(req, res proto.Message) (code uint32)),
	}, nil
}

// ListenAndServe is used to listen to TCP connection.
func ListenAndServe(addr string, handler Handler) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handler.ServeTCP(conn)
	}
}
