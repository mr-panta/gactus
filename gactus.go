package gactus

import (
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
}

// NewService [TOWRITE]
func NewService(coreAddr, tcpAddr string, coreConnPoolSize int) Service {
	return newDefaultService(coreAddr, tcpAddr, coreConnPoolSize)
}

// Processor [TOWRITE]
type Processor interface {
	GetHTTPConfig() (httpConfig *pb.HttpConfig)
	GetCommand() (command string)
	Process(req, res interface{}) (code uint32)
}
