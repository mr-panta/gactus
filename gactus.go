package gactus

import (
	"context"

	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/proto"
)

// Core [TOWRITE]
type Core interface {
	Start()
	Wait()
}

// Service [TOWRITE]
type Service interface {
	Start()
	Wait()
	RegisterService(processors []*Processor) (err error)
	SendRequest(ctx context.Context, command string, req, res proto.Message) (code uint32)
}

// Processor [TOWRITE]
type Processor struct {
	Command        string
	Req            proto.Message
	Res            proto.Message
	HTTPConfig     *pb.HttpConfig
	HTTPMiddleware func(ctx context.Context, header map[string]string, req, res proto.Message)
	Process        func(ctx context.Context, req, res proto.Message) (code uint32)
}
