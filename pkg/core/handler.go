package core

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	gactus "github.com/mr-panta/gactus/proto"
	pb "github.com/mr-panta/gactus/proto"
	"github.com/mr-panta/go-logger"
)

const (
	contentTypeJSON               = "application/json"
	contentTypeFormData           = "multipart/form-data"
	contentTypeXWWWFormURLencoded = "application/x-www-form-urlencoded"
)

// Handler [TOWRITE]
type Handler struct {
	serviceManager *serviceManager
}

// NewHandler [TOWRITE]
func NewHandler() *Handler {
	return &Handler{
		serviceManager: newServiceManager(),
	}
}

// ServeHTTP is used to implement http.Handler,
// get HTTP request and send back HTTP response.
func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	ctx, logID := generateLogID(ctx, req.Method, req.URL.Path)
	contentType, err := convertContentType(req.Header)
	if err != nil {
		logger.Errorf(ctx, "cannot convert content-type, err=%v", err)
	}

	// Get command from method and path
	command, exists := h.serviceManager.getCommand(req.Method, req.URL.Path)
	if !exists {
		logger.Errorf(ctx, "%s:%s route not found", req.Method, req.URL.Path)
		return
	}

	// Get body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Errorf(ctx, "cannot read body, err=%v", err)
		return
	}

	// Setup gactus request
	gactusReq := &gactus.Request{
		Command:     command,
		LogId:       logID,
		ContentType: contentType,
		Body:        body,
		IsProto:     false,
	}
	data, _ := proto.Marshal(gactusReq)

	// Send data to TCP
	service, exists := h.serviceManager.getServiceConn(command)
	if !exists {
		logger.Errorf(ctx, "%s command not found", command)
		return
	}
	jsonResponse, err := service.Send(data)
	if err != nil {
		logger.Errorf(ctx, "cannot send data to service, err=%v", err)
	}

	// Send response back
	res.Header().Set("Content-Type", contentTypeJSON)
	_, _ = res.Write(jsonResponse)
}

func generateLogID(ctx context.Context, method, path string) (coveredCTX context.Context, logID string) {
	if ctx == nil {
		ctx = context.Background()
	}
	coveredCTX = logger.GetContextWithLogID(ctx, fmt.Sprintf("%s_%s", method, path))
	return coveredCTX, logger.GetLogID(coveredCTX)
}

func convertContentType(header http.Header) (contentType gactus.Constant_ContentType, err error) {
	cttType := header.Get("Content-Type")
	cttTypes := strings.Split(cttType, ";")
	if len(cttTypes) == 0 {
		contentType = gactus.Constant_CONTENT_TYPE_NULL
		err = errors.New("content-type empty")
	} else {
		switch cttTypes[0] {
		case contentTypeJSON:
			contentType = gactus.Constant_CONTENT_TYPE_JSON
		case contentTypeFormData:
			contentType = gactus.Constant_CONTENT_TYPE_FORM_DATA
		case contentTypeXWWWFormURLencoded:
			contentType = gactus.Constant_CONTENT_TYPE_X_WWW_FORM_URLENCODED
		}
	}
	return
}

// ServeTCP is used to implement tcp.Handler
// and provides TCP connection.
func (h *Handler) ServeTCP(conn net.Conn) {
	ctx := context.Background()
	logger.Debugf(ctx, "new TCP connection is created")
	for {
		// receive data length
		dataSize := make([]byte, 4)
		_, err := conn.Read(dataSize)
		if err == io.EOF {
			logger.Debugf(ctx, "TCP connection is closed")
			return
		}
		if err != nil {
			logger.Errorf(ctx, err.Error())
			return
		}
		// receive data
		data := make([]byte, binary.LittleEndian.Uint32(dataSize))
		_, err = conn.Read(data)
		if err != nil {
			logger.Errorf(ctx, err.Error())
			return
		}
		wrappedReq := &pb.Request{}
		err = proto.Unmarshal(data, wrappedReq)
		if err != nil {
			logger.Errorf(ctx, err.Error())
		}
		reqCtx := logger.GetContextWithNoSubfixLogID(ctx, wrappedReq.LogId)
		// TODO: handle register data from service
		// TODO: handle request from service
		req := &pb.RegisterProcessorsRequest{}
		_ = proto.Unmarshal(wrappedReq.Body, req)
		logger.Infof(reqCtx, "%v", req)

	}
}
