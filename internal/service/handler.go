package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus/internal/config"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
	"github.com/mr-panta/go-tcpclient"
)

type handler struct {
	coreClient          tcpclient.Client
	commandProcessorMap map[string]*Processor
	commandToAddrsMap   map[string][]string
	addrToClient        map[string]tcpclient.Client
	minConns            int
	maxConns            int
	idleConnTimeout     time.Duration
	waitConnTimeout     time.Duration
	clearPeriod         time.Duration
}

// SetProcess [TOWRITE]
func (h *handler) SetProcessor(command string, processor *Processor) {
	h.commandProcessorMap[command] = processor
}

// SendCoreRequest is used to send request in bytes form to core server
func (h *handler) SendCoreRequest(logID, command string, req, res proto.Message) (code uint32, err error) {
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
func (h *handler) ServeTCP(conn net.Conn) {
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
			reqCtx := logger.GetContextWithNoSubfixLogID(ctx, wrappedReq.LogId)
			// Process reserved commands
			wrappedRes, err = h.processReservedCommands(reqCtx, wrappedReq)
			if err != nil {
				logger.Errorf(reqCtx, err.Error())
			}
			if wrappedRes != nil {
				return proto.Marshal(wrappedRes)
			}
			// Process general commands
			wrappedRes, err = h.handleRequest(reqCtx, wrappedReq)
			if err != nil {
				logger.Errorf(reqCtx, err.Error())
			}
			return proto.Marshal(wrappedRes)
		})

		if err != nil {
			logger.Warnf(ctx, "tcp connection is closed by error[%v]", err)
			return
		}
	}
}

func (h *handler) processReservedCommands(ctx context.Context, wrappedReq *pb.Request) (
	wrappedRes *pb.Response, err error) {

	switch wrappedReq.Command {
	case config.CMDServiceUpdateRegistries:
		wrappedRes, err = h.updateRegistries(ctx, wrappedReq)
	case config.CMDServiceHealthCheck:
		wrappedRes, err = h.healthCheck(ctx, wrappedReq)
	default:
		wrappedRes = nil
	}
	return wrappedRes, err
}

func (h *handler) updateRegistries(ctx context.Context, wrappedReq *pb.Request) (
	wrappedRes *pb.Response, err error) {

	logger.Debugf(ctx, "start updating registries")
	req := &pb.UpdateRegistriesRequest{}
	res := &pb.UpdateRegistriesResponse{}
	err = proto.Unmarshal(wrappedReq.Body, req)
	if err != nil {
		return nil, err
	}
	newAddrMap := make(map[string]bool)
	oldAddrMap := make(map[string]bool)
	h.commandToAddrsMap = make(map[string][]string)
	for _, pair := range req.Pairs {
		if _, exists := h.addrToClient[pair.Address]; !exists {
			newAddrMap[pair.Address] = true
		} else {
			oldAddrMap[pair.Address] = true
		}
		h.commandToAddrsMap[pair.Command] = append(h.commandToAddrsMap[pair.Command], pair.Address)
	}
	for addr, client := range h.addrToClient {
		if exists := oldAddrMap[addr]; !exists {
			// unused address
			delete(h.addrToClient, addr)
			client.Close()
		}
	}
	wrappedRes = &pb.Response{Code: uint32(pb.Constant_RESPONSE_OK)}
	for addr := range newAddrMap {
		client, err := tcpclient.NewClient(
			addr,
			h.minConns,
			h.maxConns,
			h.idleConnTimeout,
			h.waitConnTimeout,
			h.clearPeriod,
		)
		if err != nil {
			wrappedRes.Code = uint32(pb.Constant_RESPONSE_CREATE_CLIENT_FAILED)
			res.DebugMessage += fmt.Sprintf("[cannot tcp connection to service[%s]: %v]", addr, err)
		} else {
			h.addrToClient[addr] = client
		}
	}
	wrappedRes.Body, err = proto.Marshal(res)
	if err != nil {
		return nil, err
	}
	logger.Debugf(ctx, "successfully update registries")
	return wrappedRes, nil
}

func (h *handler) healthCheck(ctx context.Context, wrappedReq *pb.Request) (
	wrappedRes *pb.Response, err error) {

	logger.Debugf(ctx, "start health checking")
	wrappedRes = &pb.Response{}
	req := &pb.HealthCheckRequest{}
	res := &pb.HealthCheckResponse{}
	err = proto.Unmarshal(wrappedReq.Body, req)
	if err != nil {
		return nil, err
	}
	wrappedRes.Body, err = proto.Marshal(res)
	if err != nil {
		return nil, err
	}
	logger.Debugf(ctx, "successfully update registries")
	return wrappedRes, nil
}

func (h *handler) handleRequest(ctx context.Context, wrappedReq *pb.Request) (
	wrappedRes *pb.Response, err error) {

	wrappedRes = &pb.Response{Code: uint32(pb.Constant_RESPONSE_OK)}
	processor, exists := h.commandProcessorMap[wrappedReq.Command]
	if !exists {
		wrappedRes.Code = uint32(pb.Constant_RESPONSE_COMMAND_NOT_FOUND)
		return wrappedRes, nil
	}
	req := proto.Clone(processor.Req)
	res := proto.Clone(processor.Res)
	if wrappedReq.IsProto {
		err = proto.Unmarshal(wrappedReq.Body, req)
	} else {
		err = json.Unmarshal(wrappedReq.Body, req)
	}
	if err != nil {
		wrappedRes.Code = uint32(pb.Constant_RESPONSE_ERROR_UNPACK_REQUEST)
		return wrappedRes, err
	}
	wrappedRes.Code = processor.Process(ctx, req, res)
	if wrappedReq.IsProto {
		wrappedRes.Body, err = proto.Marshal(res)
	} else {
		wrappedRes.Body, err = json.Marshal(res)
	}
	if err != nil {
		wrappedRes.Code = uint32(pb.Constant_RESPONSE_ERROR_SETUP_RESPONSE)
	}
	return wrappedRes, err
}
