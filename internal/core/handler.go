package core

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	bd "github.com/mr-panta/gactus/internal/body"
	"github.com/mr-panta/gactus/internal/config"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
	"github.com/mr-panta/go-tcpclient"
)

// Handler [TOWRITE]
type handler struct {
	serviceManager *serviceManager
}

// ServeHTTP is used to implement http.Handler,
// get HTTP request and send back HTTP response.
func (h handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	ctx, logID := generateLogID(ctx, req.Method, req.URL.Path)

	body := []byte{}
	statusCode := http.StatusOK
	wrappedRes := &pb.Response{}
	var err error

	defer func() {
		// Send response back
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(statusCode)
		if err != nil {
			logger.Errorf(ctx, err.Error())
			body, _ = bd.Marshal(&pb.Response{
				Code:         wrappedRes.Code,
				DebugMessage: err.Error(),
			})
		}
		_, _ = res.Write(body)
	}()

	// Get content type
	contentType, rawContentType, err := bd.GetContentTypeValue(req.Header)
	if err != nil {
		err = fmt.Errorf("cannot convert content-type, err=%v", err)
		statusCode = http.StatusInternalServerError
		return
	}

	// Get command from method and path
	command, exists := h.serviceManager.getCommand(req.Method, req.URL.Path)
	if !exists {
		err = fmt.Errorf("%s:%s route not found", req.Method, req.URL.Path)
		statusCode = http.StatusNotFound
		return
	}

	// Get header
	header := make(map[string]string)
	for key, values := range req.Header {
		if len(values) > 0 {
			header[key] = values[0]
		}
	}

	// Get query
	query := make(map[string]string)
	for key, values := range req.URL.Query() {
		if len(values) > 0 {
			query[key] = values[0]
		}
	}

	// Get body
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("cannot read body, err=%v", err)
		statusCode = http.StatusInternalServerError
		return
	}

	// Get sender address
	httpAddr := strings.Split(req.RemoteAddr, ":")[0]

	// Setup gactus request
	wrappedReq := &pb.Request{
		HttpAddress:    httpAddr,
		Command:        command,
		LogId:          logID,
		ContentType:    contentType,
		RawContentType: rawContentType,
		Header:         header,
		Query:          query,
		Body:           body,
		IsProto:        false,
	}
	data, _ := proto.Marshal(wrappedReq)

	// Send data to TCP
	serviceClient, exists := h.serviceManager.getServiceConn(command)
	if !exists {
		err = fmt.Errorf("command not found, command=%v", command)
		statusCode = http.StatusNotFound
		return
	}
	output, err := serviceClient.Send(data)
	if err != nil {
		err = fmt.Errorf("cannot send data to service, err=%v", err)
		h.serviceManager.abandonService(serviceClient.GetHostAddr())
		statusCode = http.StatusInternalServerError
		return
	}

	// Unwrap response
	_ = proto.Unmarshal(output, wrappedRes)
	body = wrappedRes.Body
	if wrappedRes.Code != uint32(pb.Constant_RESPONSE_OK) {
		err = fmt.Errorf(wrappedRes.DebugMessage)
		statusCode = http.StatusInternalServerError
		return
	}
}

func generateLogID(ctx context.Context, method, path string) (coveredCTX context.Context, logID string) {
	if ctx == nil {
		ctx = context.Background()
	}
	coveredCTX = logger.GetContextWithLogID(ctx, fmt.Sprintf("%s_%s", method, path))
	return coveredCTX, logger.GetLogID(coveredCTX)
}

// ServeTCP is used to implement tcp.Handler
// and provides TCP connection.
func (h handler) ServeTCP(conn net.Conn) {
	ctx := logger.GetContextWithLogID(context.Background(), conn.RemoteAddr().String())
	logger.Debugf(ctx, "new tcp connection is created")
	for {
		err := tcpclient.Reader(conn, func(input []byte) ([]byte, error) {
			wrappedReq := &pb.Request{}
			wrappedRes := &pb.Response{}
			err := proto.Unmarshal(input, wrappedReq)
			if err != nil {
				return nil, err
			}
			reqCtx := logger.GetContextWithNoSubfixLogID(ctx, wrappedReq.LogId)

			// Find core server command
			switch wrappedReq.Command {
			case config.CMDCoreRegisterService:
				// Clean service registries
				h.serviceManager.startServiceDoctor(false)
				// Register service
				wrappedRes, err = h.serviceManager.registerService(reqCtx, wrappedReq)
				if err != nil {
					logger.Errorf(ctx, "cannot register service: error[%v]", err)
				}
				bcErr := h.serviceManager.broadcastProcessorRegistries(reqCtx)
				if bcErr != nil {
					logger.Errorf(ctx, "cannot broadcast processor registries: error[%v]", bcErr)
				}
			default:
				wrappedRes.Code = uint32(pb.Constant_RESPONSE_COMMAND_NOT_FOUND)
			}
			if err != nil {
				return nil, err
			}

			return proto.Marshal(wrappedRes)
		})

		if err != nil {
			logger.Warnf(ctx, "tcp connection is closed by error[%v]", err)
			return
		}
	}
}
