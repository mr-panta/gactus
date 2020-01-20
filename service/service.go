package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus/body"
	bd "github.com/mr-panta/gactus/body"
	"github.com/mr-panta/gactus/config"
	pb "github.com/mr-panta/gactus/proto"
	logger "github.com/mr-panta/go-logger"
	rpcpool "github.com/mr-panta/rpc-pool"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Service struct {
	lock                *sync.RWMutex
	commandProcessorMap map[string]*Processor
	commandAddressesMap map[string][]string
	addressClientMap    map[string]rpcpool.RPCPool
}

type Processor struct {
	Req            proto.Message
	Res            proto.Message
	HTTPConfig     *pb.HttpConfig
	HTTPMiddleware func(ctx context.Context, header, query map[string]string, req, res proto.Message) error
	Process        func(ctx context.Context, req, res proto.Message) error
}

func NewService() *Service {
	return &Service{
		lock:                &sync.RWMutex{},
		commandProcessorMap: make(map[string]*Processor),
		commandAddressesMap: make(map[string][]string),
		addressClientMap:    make(map[string]rpcpool.RPCPool),
	}
}

func (s *Service) getClientByCommand(command string) (client rpcpool.RPCPool, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	addrs := s.commandAddressesMap[command]
	if len(addrs) == 0 {
		return nil, errors.New("command not found")
	}
	addr := addrs[rand.Int()%len(addrs)]
	client, exists := s.addressClientMap[addr]
	if !exists {
		return nil, errors.New("address not found")
	}
	return client, nil
}

func (s *Service) AddProcessor(command string, processor *Processor) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.commandProcessorMap[command] = processor
}

func (s *Service) GetProcessor(command string) (processor *Processor, exists bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	p, exists := s.commandProcessorMap[command]
	return p, exists
}

func (s *Service) Receive(wrappedReq *pb.Request, wrappedRes *pb.Response) (err error) {
	if wrappedReq == nil {
		return fmt.Errorf("request nil")
	}
	if wrappedRes == nil {
		return fmt.Errorf("response nil")
	}
	ctx := logger.GetContextWithLogID(context.Background(), wrappedReq.LogId)
	processor, exists := s.GetProcessor(wrappedReq.Command)
	if !exists {
		return fmt.Errorf("command not found, command=%s", wrappedReq.Command)
	}
	req := proto.Clone(processor.Req)
	res := proto.Clone(processor.Res)
	if wrappedReq.IsProto {
		err = proto.Unmarshal(wrappedReq.Body, req)
	} else if wrappedReq.ContentType != pb.Constant_CONTENT_TYPE_UNKNOWN {
		err = bd.Unmarshal(wrappedReq, req)
	}
	if err != nil {
		return fmt.Errorf("cannot unpack request, err=%v", err)
	}
	if !wrappedReq.IsProto && processor.HTTPMiddleware != nil {
		err = processor.HTTPMiddleware(ctx, wrappedReq.Header, wrappedReq.Query, req, res)
		if err != nil {
			return err
		}
	}
	err = processor.Process(ctx, req, res)
	if err != nil {
		return err
	}
	if wrappedReq.IsProto {
		wrappedRes.Body, err = proto.Marshal(res)
	} else {
		wrappedRes.Body, err = body.Marshal(res)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SendRequestWithAddress(ctx context.Context, address string, command string, req proto.Message, res proto.Message) error {
	body, err := proto.Marshal(req)
	if err != nil {
		return err
	}
	wrappedReq := &pb.Request{
		LogId:   logger.GetLogID(ctx),
		Command: command,
		IsProto: true,
		Body:    body,
	}
	wrappedRes := &pb.Response{}
	err = s.SendWrappedRequest(address, wrappedReq, wrappedRes)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(wrappedRes.Body, res)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SendRequest(ctx context.Context, command string, req proto.Message, res proto.Message) error {
	return s.SendRequestWithAddress(ctx, "", command, req, res)
}

func (s *Service) SendWrappedRequest(address string, wrappedReq *pb.Request, wrappedRes *pb.Response) (err error) {
	if wrappedReq == nil {
		return errors.New("wrapped request empty")
	}
	if wrappedRes == nil {
		return errors.New("wrapped response empty")
	}
	var client rpcpool.RPCPool
	if address == "" {
		client, err = s.getClientByCommand(wrappedReq.Command)
		if err != nil {
			return err
		}
	} else {
		var exists bool
		client, exists = s.addressClientMap[address]
		if !exists {
			return fmt.Errorf("cannot find rpc client, address=%s", address)
		}
	}
	err = client.Call(config.ReceiverMethodName, wrappedReq, wrappedRes)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SetAddressCommands(address string, commands []string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, exists := s.addressClientMap[address]
	if !exists {
		client, err := rpcpool.NewRPCPool(
			address,
			config.MinConns,
			config.MaxConns,
			config.IdleConnTimeout,
			config.WaitConnTimeout,
			config.ClearPeriod,
		)
		if err != nil {
			return err
		}
		s.addressClientMap[address] = client
	}
	for _, command := range commands {
		addrs := s.commandAddressesMap[command]
		exists = false
		for _, addr := range addrs {
			if addr == address {
				exists = true
				break
			}
		}
		if !exists {
			addrs = append(addrs, address)
		}
		s.commandAddressesMap[command] = addrs
	}
	return nil
}

func (s *Service) GetAddressCommandSet() []*pb.AddressCommandSet {
	s.lock.RLock()
	defer s.lock.RUnlock()
	addrCmds := []*pb.AddressCommandSet{}
	addrCmdsMap := make(map[string][]string)
	for cmd, addrs := range s.commandAddressesMap {
		for _, addr := range addrs {
			addrCmdsMap[addr] = append(addrCmdsMap[addr], cmd)
		}
	}
	for addr, cmds := range addrCmdsMap {
		addrCmds = append(addrCmds, &pb.AddressCommandSet{
			Address:   addr,
			Commmands: cmds,
		})
	}
	return addrCmds
}
