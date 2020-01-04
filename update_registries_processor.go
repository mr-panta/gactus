package gactus

import (
	"context"
	"errors"

	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/proto"
)

func (gs *gactusService) ProcessUpdateRegistries(ctx context.Context, request proto.Message, response proto.Message) (err error) {
	req, ok := request.(*pb.UpdateRegistriesRequest)
	if !ok {
		return errors.New("cannot assert request object")
	}
	addrCmds := req.AddressCommands
	for _, addrCmd := range addrCmds {
		err = gs.service.SetAddressCommands(addrCmd.Address, addrCmd.Commmands)
		if err != nil {
			return err
		}
	}
	return nil
}
