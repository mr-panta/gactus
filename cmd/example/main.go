package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus"
	"github.com/mr-panta/gactus/cmd/example/echo"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
)

func main() {
	core := gactus.NewGactusCore(8000, 3000, "hi")
	core.Start()

	s := gactus.NewGactusService("s", 4000, "localhost:3000", "hi")
	s.Start()
	logger.Infof(context.Background(), "gactus core started on http port=%d, tcp port=%d, with access-key=%s")
	err := s.RegisterProcessors([]*gactus.Processor{
		{
			Command: "s.do",
			Req:     &echo.Request{},
			Res:     &echo.Response{},
			Process: func(ctx context.Context, request, response proto.Message) error {
				fmt.Println("s.do called")
				req := request.(*echo.Request)
				res := response.(*echo.Response)
				res.Message = req.Message + ", Okay"
				return nil
			},
		},
		{
			Command: "s.redirect",
			Req:     &echo.Request{},
			Res:     &echo.Response{},
			HTTPConfig: &pb.HttpConfig{
				Method: pb.Constant_HTTP_METHOD_GET,
				Path:   "/redirect",
			},
			HTTPMiddleware: func(ctx context.Context, header, query map[string]string, request, response proto.Message) error {
				req := request.(*echo.Request)
				req.Message = query["message"]
				return nil
			},
			Process: func(ctx context.Context, request, response proto.Message) error {
				fmt.Println("s.redirect called")
				return s.SendRequest(ctx, "s2.redirect", request, response)
			},
		},
	})
	if err != nil {
		logger.Fatalf(context.Background(), err.Error())
	}

	s2 := gactus.NewGactusService("s2", 5000, "localhost:3000", "hi")
	s2.Start()
	err = s2.RegisterProcessors([]*gactus.Processor{
		{
			Command: "s2.redirect",
			Req:     &echo.Request{},
			Res:     &echo.Response{},
			HTTPConfig: &pb.HttpConfig{
				Method: pb.Constant_HTTP_METHOD_GET,
				Path:   "/echo",
			},
			HTTPMiddleware: func(ctx context.Context, header, query map[string]string, request, response proto.Message) error {
				req := request.(*echo.Request)
				req.Message = query["message"]
				return nil
			},
			Process: func(ctx context.Context, request, response proto.Message) error {
				fmt.Println("s2.redirect called")
				return s2.SendRequest(ctx, "s.do", request, response)
			},
		},
	})
	if err != nil {
		logger.Fatalf(context.Background(), err.Error())
	}

	s2.Wait()
	s.Wait()
	core.Wait()
}
