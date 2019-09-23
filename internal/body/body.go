package body

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	pb "github.com/mr-panta/gactus/proto"
)

// HTTPBody implements io.ReadCloser to keep HTTP body data.
type HTTPBody struct {
	i        int
	p        []byte
	isClosed bool
}

// NewHTTPBody is used to create HTTPBody.
func NewHTTPBody(p []byte) *HTTPBody {
	return &HTTPBody{p: p}
}

// Raw is used to get raw HTTP body directly without removing the data.
func (b *HTTPBody) Raw() []byte {
	return b.p
}

// Read is used to implement io.Reader.
func (b *HTTPBody) Read(p []byte) (n int, err error) {
	if b.isClosed {
		return 0, errors.New("cannot read")
	}
	for n = 0; n < len(p) && b.i < len(b.p); n++ {
		p[n] = b.p[b.i]
		b.i++
	}
	if len(b.p) == b.i {
		err = io.EOF
	}
	return
}

// Close is used to implement io.Closer.
func (b *HTTPBody) Close() error {
	b.i = 0
	b.p = nil
	b.isClosed = true
	return nil
}

// Content types in http request.
const (
	contentTypeJSON               = "application/json"
	contentTypeFormData           = "multipart/form-data"
	contentTypeXWWWFormURLencoded = "application/x-www-form-urlencoded"
)

// GetContentTypeValue is used to convert content type from http request
// to enum value.
func GetContentTypeValue(header http.Header) (contentType pb.Constant_ContentType, rawContentType string, err error) {
	rawContentType = header.Get("Content-Type")
	cttTypes := strings.Split(rawContentType, ";")
	if len(cttTypes) == 0 {
		contentType = pb.Constant_CONTENT_TYPE_UNKNOWN
		err = errors.New("content-type empty")
	} else {
		switch cttTypes[0] {
		case contentTypeJSON:
			rawContentType = ""
			contentType = pb.Constant_CONTENT_TYPE_JSON
		case contentTypeFormData:
			contentType = pb.Constant_CONTENT_TYPE_FORM_DATA
		case contentTypeXWWWFormURLencoded:
			rawContentType = ""
			contentType = pb.Constant_CONTENT_TYPE_X_WWW_FORM_URLENCODED
		}
	}
	return
}

// GetContenTypeString is used to convert enum value of content type to string.
func GetContenTypeString(contentType pb.Constant_ContentType) string {
	switch contentType {
	case pb.Constant_CONTENT_TYPE_JSON:
		return contentTypeJSON
	case pb.Constant_CONTENT_TYPE_FORM_DATA:
		return contentTypeFormData
	case pb.Constant_CONTENT_TYPE_X_WWW_FORM_URLENCODED:
		return contentTypeXWWWFormURLencoded
	}
	// default type
	return contentTypeJSON
}

// Marshal is used to marshal protobuf message to http response body
// (JSON format only).
func Marshal(msg proto.Message) (b []byte, err error) {
	return json.Marshal(msg)
}

// Unmarshal is used to unmarshal http body to protobuf message.
func Unmarshal(wrappedReq *pb.Request, msg proto.Message) (err error) {
	req := &http.Request{
		Method: http.MethodPost,
		Body:   NewHTTPBody(wrappedReq.Body),
		Header: make(map[string][]string),
	}
	switch wrappedReq.ContentType {
	case pb.Constant_CONTENT_TYPE_JSON:
		return json.Unmarshal(wrappedReq.Body, msg)
	case pb.Constant_CONTENT_TYPE_FORM_DATA:
		req.Header.Set("Content-Type", wrappedReq.RawContentType)
		err = req.ParseMultipartForm((1 << 20) * 24) // TODO: get this value by config
	case pb.Constant_CONTENT_TYPE_X_WWW_FORM_URLENCODED:
		ctt := GetContenTypeString(wrappedReq.ContentType)
		req.Header.Set("Content-Type", ctt)
		err = req.ParseForm()
	default:
		err = errors.New("content type is not supported")
	}
	if err != nil {
		return err
	}
	// TEST
	fmt.Println("A", req.FormValue("a"))
	fmt.Println("B", req.FormValue("b"))
	if wrappedReq.ContentType == pb.Constant_CONTENT_TYPE_FORM_DATA {
		_, fHeader, err := req.FormFile("file")
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("File", fHeader.Filename)
		}
	}
	// TODO: implement unmarshaler
	return nil
}
