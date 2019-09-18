package service

import (
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/tcpclient"
)

// Handler [TOWRITE]
type defaultHandler struct {
	coreClient        tcpclient.Client
	commandProcessMap map[string]func(req, res proto.Message) (code uint32)
}

// NewHandler [TOWRITE]
func NewHandler(coreAddr string, minConns, maxConns, idleConnTimeout, waitConnTimeout, clearPeriod int) (hanlder Handler, err error) {
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
	return &defaultHandler{
		coreClient:        coreClient,
		commandProcessMap: make(map[string]func(req, res proto.Message) (code uint32)),
	}, nil
}

// SetProcess [TOWRITE]
func (h *defaultHandler) SetProcess(command string, process func(req, res proto.Message) (code uint32)) {
	h.commandProcessMap[command] = process
}

func (h *defaultHandler) handleRequest(command string, req, res proto.Message) (code uint32) {
	process, exists := h.commandProcessMap[command]
	if !exists {
		return uint32(pb.Constant_RESPONSE_PROCESS_NOT_FOUND)
	}
	return process(req, res)
}

// SendCoreRequest is used to send request in bytes form to core server
func (h *defaultHandler) SendCoreRequest(logID, command string, req, res proto.Message) (code uint32, err error) {
	body, err := proto.Marshal(req)
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	resData, err := proto.Marshal(&pb.Request{
		LogId:   logID,
		Command: command,
		IsProto: true,
		Body:    body,
	})
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	data, err := h.coreClient.Send(resData)
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	err = proto.Unmarshal(data, res)
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	return uint32(pb.Constant_RESPONSE_OK), nil
}

// ServeTCP is used to implement tcp.Handler
// and provides TCP connection.
func (h *defaultHandler) ServeTCP(conn net.Conn) {}
