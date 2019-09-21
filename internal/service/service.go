package service

import (
	"context"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-tcpclient"
)

// Processor [TOWRITE]
type Processor struct {
	Req     proto.Message
	Res     proto.Message
	Process func(ctx context.Context, req, res proto.Message) (code uint32)
}

// Handler [TOWRITE]
type Handler interface {
	SendCoreRequest(logID, command string, req, res proto.Message) (code uint32, err error)
	ServeTCP(conn net.Conn)
	SetProcessor(command string, processor *Processor)
	SetTCPAddr(addr string)
}

// NewHandler [TOWRITE]
func NewHandler(coreAddr string, minConns, maxConns, idleConnTimeout, waitConnTimeout, clearPeriod int) (Handler, error) {
	coreClient, err := tcpclient.NewClient(
		coreAddr,
		0,
		1,
		time.Duration(idleConnTimeout)*time.Millisecond,
		time.Duration(waitConnTimeout)*time.Millisecond,
		time.Duration(clearPeriod)*time.Millisecond,
	)
	if err != nil {
		return nil, err
	}
	return &handler{
		coreClient:          coreClient,
		commandProcessorMap: make(map[string]*Processor),
		commandToAddrsMap:   make(map[string][]string),
		addrToClientMap:     make(map[string]tcpclient.Client),
		addrToConnConfigMap: make(map[string]*pb.ConnectionConfig),
		minConns:            minConns,
		maxConns:            maxConns,
		idleConnTimeout:     time.Duration(idleConnTimeout) * time.Millisecond,
		waitConnTimeout:     time.Duration(waitConnTimeout) * time.Millisecond,
		clearPeriod:         time.Duration(clearPeriod) * time.Millisecond,
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
