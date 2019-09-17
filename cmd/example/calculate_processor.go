package main

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/cmd/example/proto"
	gtpb "github.com/mr-panta/gactus/proto"
)

type calculateProcessor struct{}

func (p *calculateProcessor) GetCommand() (command string) {
	return "example.calculate"
}

func (p *calculateProcessor) GetRequest() (req proto.Message) {
	return &pb.CalculateRequest{}
}

func (p *calculateProcessor) GetResponse() (req proto.Message) {
	return &pb.CalculateResponse{}
}

func (p *calculateProcessor) GetHTTPConfig() (httpConfig *gtpb.HttpConfig) {
	return &gtpb.HttpConfig{
		Method: gtpb.Constant_HTTP_METHOD_POST,
		Path:   "/calculate",
	}
}

func (p *calculateProcessor) Process(req, res proto.Message) (code uint32) {
	request := req.(*pb.CalculateRequest)
	response := res.(*pb.CalculateResponse)
	response.C = request.A + request.B
	return uint32(gtpb.Constant_RESPONSE_OK)
}
