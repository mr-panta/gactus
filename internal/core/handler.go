package core

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus/internal/config"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
	"github.com/mr-panta/go-tcpclient"
)

const (
	contentTypeJSON               = "application/json"
	contentTypeFormData           = "multipart/form-data"
	contentTypeXWWWFormURLencoded = "application/x-www-form-urlencoded"
)

// Handler [TOWRITE]
type handler struct {
	serviceManager *serviceManager
}

// ServeHTTP is used to implement http.Handler,
// get HTTP request and send back HTTP response.
func (h handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	body := []byte{}
	statusCode := http.StatusOK
	defer func() {
		// Send response back
		if statusCode == http.StatusOK {
			res.Header().Set("Content-Type", contentTypeJSON)
		}
		res.WriteHeader(statusCode)
		_, _ = res.Write(body)
	}()
	ctx := context.Background()
	ctx, logID := generateLogID(ctx, req.Method, req.URL.Path)
	contentType, err := convertContentType(req.Header)
	if err != nil {
		logger.Errorf(ctx, "cannot convert content-type, err=%v", err)
		statusCode = http.StatusInternalServerError
		body = []byte(fmt.Sprintf("%d %s", statusCode, config.ErrorServiceNotAvailable))
		return
	}

	// Get command from method and path
	command, exists := h.serviceManager.getCommand(req.Method, req.URL.Path)
	if !exists {
		logger.Errorf(ctx, "%s:%s route not found", req.Method, req.URL.Path)
		statusCode = http.StatusNotFound
		body = []byte(fmt.Sprintf("%d %s", statusCode, config.ErrorNotFound))
		return
	}

	// Get body
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Errorf(ctx, "cannot read body, err=%v", err)
		statusCode = http.StatusInternalServerError
		body = []byte(fmt.Sprintf("%d %s", statusCode, config.ErrorServiceNotAvailable))
		return
	}

	// Setup gactus request
	wrappedReq := &pb.Request{
		Command:     command,
		LogId:       logID,
		ContentType: contentType,
		Body:        body,
		IsProto:     false,
	}
	data, _ := proto.Marshal(wrappedReq)

	// Send data to TCP
	serviceClient, exists := h.serviceManager.getServiceConn(command)
	if !exists {
		logger.Errorf(ctx, "%s command not found", command)
		statusCode = http.StatusNotFound
		body = []byte(fmt.Sprintf("%d %s", statusCode, config.ErrorNotFound))
		return
	}
	output, err := serviceClient.Send(data)
	if err != nil {
		logger.Errorf(ctx, "cannot send data to service, err=%v", err)
		h.serviceManager.abandonService(serviceClient.GetHostAddr())
		statusCode = http.StatusInternalServerError
		body = []byte(fmt.Sprintf("%d %s", statusCode, config.ErrorServiceNotAvailable))
		return
	}
	h.serviceManager.addrToActiveTimeMap[serviceClient.GetHostAddr()] = time.Now()

	// Unwrap response
	wrappedRes := &pb.Response{}
	_ = proto.Unmarshal(output, wrappedRes)
	body = wrappedRes.Body
	if wrappedRes.Code != uint32(pb.Constant_RESPONSE_OK) {
		errorMessage := fmt.Sprintf("internal error code: %d", wrappedRes.Code)
		logger.Errorf(ctx, errorMessage)
		statusCode = http.StatusInternalServerError
		body = []byte(errorMessage)
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

func convertContentType(header http.Header) (contentType pb.Constant_ContentType, err error) {
	cttType := header.Get("Content-Type")
	cttTypes := strings.Split(cttType, ";")
	if len(cttTypes) == 0 {
		contentType = pb.Constant_CONTENT_TYPE_UNKNOWN
		err = errors.New("content-type empty")
	} else {
		switch cttTypes[0] {
		case contentTypeJSON:
			contentType = pb.Constant_CONTENT_TYPE_JSON
		case contentTypeFormData:
			contentType = pb.Constant_CONTENT_TYPE_FORM_DATA
		case contentTypeXWWWFormURLencoded:
			contentType = pb.Constant_CONTENT_TYPE_X_WWW_FORM_URLENCODED
		}
	}
	return
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
			case config.CMDCoreRegisterProcessors:
				wrappedRes, err = h.serviceManager.registerProcessors(reqCtx, wrappedReq)
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
			logger.Errorf(ctx, "tcp connection is closed by error[%v]", err)
			return
		}
	}
}
