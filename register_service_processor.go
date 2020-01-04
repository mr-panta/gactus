package gactus

import (
	"context"
	"errors"
	"time"

	"github.com/mr-panta/gactus/config"

	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
	rpcpool "github.com/mr-panta/rpc-pool"
)

func verifyServiceAddresses(ctx context.Context, addrs []string) (addr string, err error) {
	for _, addr := range addrs {
		client, err := rpcpool.NewRPCPool(addr, 1, 1, 0, 0, time.Second)
		if err == nil {
			client.Close()
			return addr, nil
		}
		logger.Errorf(ctx, err.Error())
	}
	return "", errors.New("cannot verify service addresses")
}

func (gc *gactusCore) UpdateAllServices(ctx context.Context) error {
	addrCmds := gc.service.GetAddressCommandSet()
	req := &pb.UpdateRegistriesRequest{
		AddressCommands: addrCmds,
	}
	res := &pb.UpdateRegistriesResponse{}
	for _, addrCmd := range addrCmds {
		err := gc.service.SendRequestWithAddress(ctx, addrCmd.Address, config.CMDServiceUpdateRegistries, req, res)
		if err != nil {
			logger.Warnf(ctx, "cannot update send updated registries, err=%v", err)
		}
	}
	return nil
}

func (gc *gactusCore) ProcessRegisterService(ctx context.Context, request proto.Message, response proto.Message) (err error) {
	req, ok := request.(*pb.RegisterServiceRequest)
	if !ok {
		return errors.New("cannot assert request object")
	}
	res, ok := response.(*pb.RegisterServiceResponse)
	if !ok {
		return errors.New("cannot assert response object")
	}
	addr, err := verifyServiceAddresses(ctx, req.Addresses)
	if err != nil {
		return err
	}
	commands := []string{}
	for _, processor := range req.ProcessorRegistries {
		if processor.HttpConfig != nil {
			method := getMethodString(processor.HttpConfig.Method)
			route := getRoute(method, processor.HttpConfig.Path)
			gc.setCommandByRoute(route, processor.Command)
		}
		commands = append(commands, processor.Command)
	}
	err = gc.service.SetAddressCommands(addr, commands)
	if err != nil {
		return err
	}
	err = gc.UpdateAllServices(ctx)
	if err != nil {
		return err
	}
	res.Address = addr
	return nil
}
