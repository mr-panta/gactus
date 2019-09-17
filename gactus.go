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

// NewCore [TOWRITE]
func NewCore(httpAddr, tcpAddr string) Core {
	return newDefaultCore(httpAddr, tcpAddr)
}

// Service [TOWRITE]
type Service interface {
	Start()
	Wait()
	RegisterProcessors(processors []Processor) (err error)
	SendRequest(ctx context.Context, command string, req, res proto.Message) (code uint32)
}

// NewService [TOWRITE]
func NewService(serviceName, coreAddr, tcpAddr string, coreConnPoolSize int) Service {
	return newDefaultService(serviceName, coreAddr, tcpAddr, coreConnPoolSize)
}

// Processor [TOWRITE]
type Processor interface {
	GetCommand() (command string)
	GetRequest() (req proto.Message)
	GetResponse() (res proto.Message)
	GetHTTPConfig() (httpConfig *pb.HttpConfig)
	Process(req, res proto.Message) (code uint32)
}
