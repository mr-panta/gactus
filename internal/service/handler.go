package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	bd "github.com/mr-panta/gactus/internal/body"
	"github.com/mr-panta/gactus/internal/config"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
	"github.com/mr-panta/go-tcpclient"
)

type handler struct {
	coreClient          tcpclient.Client
	commandProcessorMap map[string]*Processor
	commandToAddrsMap   map[string][]string
	addrToClientMap     map[string]tcpclient.Client
	tcpAddr             string
	minConns            int
	maxConns            int
	idleConnTimeout     time.Duration
	waitConnTimeout     time.Duration
	clearPeriod         time.Duration
	lock                sync.Mutex
}

// SetProcess [TOWRITE]
func (h *handler) SetProcessor(command string, processor *Processor) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.commandProcessorMap[command] = processor
}

func (h *handler) setupCoreData() {
	h.commandToAddrsMap[config.CMDCoreRegisterService] = append(
		h.commandToAddrsMap[config.CMDCoreRegisterService],
		h.coreClient.GetHostAddr(),
	)
	h.addrToClientMap[h.coreClient.GetHostAddr()] = h.coreClient
}

func (h *handler) getServiceConn(command string) (service tcpclient.Client, exists bool) {
	addrs := h.commandToAddrsMap[command]
	addrsLength := len(addrs)
	if addrsLength == 0 {
		return nil, false
	}
	addr := addrs[rand.Intn(addrsLength)]
	service, exists = h.addrToClientMap[addr]
	return
}

// SendRequest is used to send request in bytes form to core server
func (h *handler) SendRequest(logID, command string, req, res proto.Message) (code uint32, err error) {
	body, err := proto.Marshal(req)
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	wrappedReq, err := proto.Marshal(&pb.Request{
		LogId:   logID,
		Command: command,
		IsProto: true,
		Body:    body,
	})
	wrappedRes := &pb.Response{}
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	serviceClient, exists := h.getServiceConn(command)
	if !exists {
		return uint32(pb.Constant_RESPONSE_COMMAND_NOT_FOUND), errors.New("command not found")
	}
	data, err := serviceClient.Send(wrappedReq)
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	err = proto.Unmarshal(data, wrappedRes)
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	err = proto.Unmarshal(wrappedRes.Body, res)
	if err != nil {
		return uint32(pb.Constant_RESPONSE_ERROR_SETUP_REQUEST), err
	}
	return uint32(wrappedRes.Code), nil
}

// ServeTCP is used to implement tcp.Handler
// and provides TCP connection.
func (h *handler) ServeTCP(conn net.Conn) {
	ctx := logger.GetContextWithLogID(context.Background(), conn.RemoteAddr().String())
	logger.Debugf(ctx, "new tcp connection is created")
	for {
		err := tcpclient.Reader(conn, func(input []byte) ([]byte, error) {
			wrappedReq := &pb.Request{}
			err := proto.Unmarshal(input, wrappedReq)
			if err != nil {
				return nil, err
			}
			reqCtx := logger.GetContextWithNoSubfixLogID(ctx, wrappedReq.LogId)
			// Process reserved commands
			wrappedRes, err := h.processReservedCommands(reqCtx, wrappedReq)
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

	h.lock.Lock()
	defer h.lock.Unlock()
	logger.Debugf(ctx, "start updating registries")
	wrappedRes = &pb.Response{}
	req := &pb.UpdateRegistriesRequest{}
	res := &pb.UpdateRegistriesResponse{}
	err = proto.Unmarshal(wrappedReq.Body, req)
	if err != nil {
		return nil, err
	}
	h.commandToAddrsMap = make(map[string][]string)
	h.setupCoreData()
	for addr := range h.addrToClientMap {
		if addr == h.coreClient.GetHostAddr() {
			// Skip core server address
			continue
		}
		client := h.addrToClientMap[addr]
		delete(h.addrToClientMap, addr)
		client.Close()
	}
	for _, addrConf := range req.AddrConfigs {
		if addrConf.Address == h.tcpAddr {
			continue
		}
		client, err := tcpclient.NewClient(
			addrConf.Address,
			int(addrConf.ConnConfig.MinConns),
			int(addrConf.ConnConfig.MaxConns),
			time.Duration(addrConf.ConnConfig.IdleConnTimeout)*time.Microsecond,
			time.Duration(addrConf.ConnConfig.WaitConnTimeout)*time.Microsecond,
			time.Duration(addrConf.ConnConfig.ClearPeriod)*time.Microsecond,
		)
		if err != nil {
			wrappedRes.Code = uint32(pb.Constant_RESPONSE_CREATE_CLIENT_FAILED)
			res.DebugMessage += fmt.Sprintf("[cannot tcp connection to service[%s]: %v]", addrConf.Address, err)
		} else {
			h.addrToClientMap[addrConf.Address] = client
		}
	}
	for _, pair := range req.CommandAddressPairs {
		if pair.Address == h.tcpAddr {
			continue
		}
		h.commandToAddrsMap[pair.Command] = append(h.commandToAddrsMap[pair.Command], pair.Address)
	}
	wrappedRes.Code = uint32(pb.Constant_RESPONSE_OK)
	wrappedRes.Body, err = proto.Marshal(res)
	if err != nil {
		return nil, err
	}
	logger.Debugf(ctx, "successfully update registries")
	return wrappedRes, nil
}

func (h *handler) healthCheck(ctx context.Context, wrappedReq *pb.Request) (
	wrappedRes *pb.Response, err error) {

	wrappedRes = &pb.Response{}
	req := &pb.HealthCheckRequest{}
	res := &pb.HealthCheckResponse{}
	err = proto.Unmarshal(wrappedReq.Body, req)
	if err != nil {
		return nil, err
	}
	if h.tcpAddr == "" {
		logger.Debugf(ctx, "assign tcp address[%s]", req.Address)
		h.tcpAddr = req.Address
	}
	wrappedRes.Body, err = proto.Marshal(res)
	if err != nil {
		return nil, err
	}
	return wrappedRes, nil
}

func (h *handler) handleRequest(ctx context.Context, wrappedReq *pb.Request) (
	wrappedRes *pb.Response, err error) {

	wrappedRes = &pb.Response{Code: uint32(pb.Constant_RESPONSE_OK)}
	defer func() {
		if err != nil {
			wrappedRes.DebugMessage = err.Error()
		}
	}()

	processor, exists := h.commandProcessorMap[wrappedReq.Command]
	if !exists {
		wrappedRes.Code = uint32(pb.Constant_RESPONSE_COMMAND_NOT_FOUND)
		return wrappedRes, nil
	}
	req := proto.Clone(processor.Req)
	res := proto.Clone(processor.Res)
	if wrappedReq.IsProto {
		err = proto.Unmarshal(wrappedReq.Body, req)
	} else if wrappedReq.ContentType != pb.Constant_CONTENT_TYPE_UNKNOWN {
		err = bd.Unmarshal(wrappedReq, req)
	}
	if err != nil {
		wrappedRes.Code = uint32(pb.Constant_RESPONSE_ERROR_UNPACK_REQUEST)
		return wrappedRes, err
	}
	if !wrappedReq.IsProto && processor.HTTPMiddleware != nil {
		wrappedRes.Code = processor.HTTPMiddleware(ctx, wrappedReq.Header, wrappedReq.Query, req, res)
	}
	if wrappedRes.Code == uint32(pb.Constant_RESPONSE_OK) {
		wrappedRes.Code = processor.Process(ctx, req, res)
	}
	if wrappedReq.IsProto {
		wrappedRes.Body, err = proto.Marshal(res)
	} else {
		wrappedRes.Body, err = bd.Marshal(res)
	}
	if err != nil {
		wrappedRes.Code = uint32(pb.Constant_RESPONSE_ERROR_SETUP_RESPONSE)
	}
	return wrappedRes, err
}
