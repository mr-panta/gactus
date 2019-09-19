package service

import (
	"context"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
	"github.com/mr-panta/go-tcpclient"
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
		return uint32(pb.Constant_RESPONSE_COMMAND_NOT_FOUND)
	}
	return process(req, res)
}

// SendCoreRequest is used to send request in bytes form to core server
func (h *defaultHandler) SendCoreRequest(logID, command string, req, res proto.Message) (code uint32, err error) {
	body, err := proto.Marshal(req)
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	reqData, err := proto.Marshal(&pb.Request{
		LogId:   logID,
		Command: command,
		IsProto: true,
		Body:    body,
	})
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	data, err := h.coreClient.Send(reqData)
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
func (h *defaultHandler) ServeTCP(conn net.Conn) {
	ctx := logger.GetContextWithLogID(context.Background(), conn.RemoteAddr().String())
	logger.Debugf(ctx, "new tcp connection is created")
	for {
		err := tcpclient.Reader(conn, func(input []byte) ([]byte, error) {
			wrappedReq := &pb.Request{}
			wrappedRes := &pb.Response{}
			err := proto.Unmarshal(input, wrappedReq)
			if err != nil {
				return nil, err
			}
			// reqCtx := logger.GetContextWithNoSubfixLogID(ctx, wrappedReq.LogId)
			// TODO: process reserve commands
			// TODO: process general commands
			_ = h.handleRequest("", nil, nil) // TODO: setup parameters
			return proto.Marshal(wrappedRes)
		})

		if err != nil {
			logger.Errorf(ctx, "tcp connection is closed by error[%v]", err)
			return
		}
	}
}