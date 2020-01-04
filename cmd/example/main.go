package main

import (
	"context"

	"github.com/mr-panta/go-logger"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus"
	"github.com/mr-panta/gactus/cmd/example/echo"
	pb "github.com/mr-panta/gactus/proto"
)

// func main() {
// 	s1 := gactus.NewGactusService(3000)
// 	s1.RegisterProcessors([]*gactus.Processor{
// 		{
// 			Command: "echo.do",
// 			Req:     &echo.Request{},
// 			Res:     &echo.Response{},
// 			Process: func(ctx context.Context, request, response proto.Message) error {
// 				req := request.(*echo.Request)
// 				res := response.(*echo.Response)
// 				res.Message = req.Message + ", Okay"
// 				return nil
// 			},
// 		},
// 	})
// 	go func() {
// 		err := s1.Start()
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}()

// 	time.Sleep(time.Second)

// 	s2 := service.NewService()
// 	err := s2.SetAddressCommands(":3000", []string{"echo.do"})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	req := &echo.Request{Message: "Hi"}
// 	res := &echo.Response{}
// 	err = s2.SendRequest(context.Background(), "echo.do", req, res)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(res.Message)
// }

func main() {
	core := gactus.NewGactusCore(8000, 3000)
	core.Start()

	s := gactus.NewGactusService("localhost:3000", 4000)
	s.Start()
	err := s.RegisterProcessors([]*gactus.Processor{
		{
			Command: "echo.do",
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
				req := request.(*echo.Request)
				res := response.(*echo.Response)
				res.Message = req.Message + ", Okay"
				return nil
			},
		},
	})
	if err != nil {
		logger.Fatalf(context.Background(), err.Error())
	}
	s.Wait()
	core.Wait()
}
