package body

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/mr-panta/gactus/config"
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

// CheckContentType is used to validate if the content type exists or not.
func CheckContentType(contentType string) bool {
	return contentType == contentTypeJSON ||
		contentType == contentTypeFormData ||
		contentType == contentTypeXWWWFormURLencoded
}

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
	if msg == nil {
		return errors.New("msg is nil")
	}
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
		return unmarshalFormData(req, msg)
	case pb.Constant_CONTENT_TYPE_X_WWW_FORM_URLENCODED:
		req.Header.Set("Content-Type", GetContenTypeString(wrappedReq.ContentType))
		return unmarshalXWWWURLEncoded(req, msg)
	}
	return errors.New("content type is not supported")
}

func unmarshalFormData(req *http.Request, msg proto.Message) (err error) {
	err = req.ParseMultipartForm((1 << 20) * config.LimitSize)
	if err != nil {
		return err
	}
	form := req.MultipartForm
	v := reflect.ValueOf(msg).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		tag, exists := getJSONTag(t.Field(i))
		if !exists {
			continue
		}
		// Map basic data type
		found, err := mapBasicDataType(v.Field(i), req.FormValue(tag))
		if err != nil {
			return err
		}
		// Map complex data type
		if !found {
			if v.Field(i).Kind() == reflect.Slice {
				elemKind := v.Field(i).Type().Elem()
				typePath := strings.Split(elemKind.String(), ".")
				if typePath[len(typePath)-1] == "GactusFile" {
					for _, file := range form.File[tag] {
						elem := reflect.New(elemKind).Elem()
						_, err = mapGactusFileType(elem, file)
						if err != nil {
							return err
						}
						v.Field(i).Set(reflect.Append(v.Field(i), elem))
					}
				} else {
					for _, value := range form.Value[tag] {
						elem := reflect.New(elemKind).Elem()
						_, err = mapBasicDataType(elem, value)
						if err != nil {
							return err
						}
						v.Field(i).Set(reflect.Append(v.Field(i), elem))
					}
				}
			} else if len(form.File[tag]) > 0 {
				_, err = mapGactusFileType(v.Field(i), form.File[tag][0])
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func unmarshalXWWWURLEncoded(req *http.Request, msg proto.Message) (err error) {
	err = req.ParseForm()
	if err != nil {
		return
	}
	v := reflect.ValueOf(msg).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		tag, exists := getJSONTag(t.Field(i))
		if !exists {
			continue
		}
		_, err := mapBasicDataType(v.Field(i), req.FormValue(tag))
		if err != nil {
			return err
		}
	}
	return nil
}

func getJSONTag(field reflect.StructField) (tag string, exists bool) {
	fullTag := field.Tag.Get("json")
	if len(fullTag) == 0 || fullTag == "-" {
		return "", false
	}
	sTag := strings.Split(fullTag, ",")
	return sTag[0], true
}

func mapGactusFileType(v reflect.Value, fileHeader *multipart.FileHeader) (match bool, err error) {
	typePath := strings.Split(v.Type().String(), ".")
	if typePath[len(typePath)-1] == "GactusFile" {

		gactusFile := reflect.New(v.Type().Elem()).Elem()
		nameField, nameExists := gactusFile.Type().FieldByName("Name")
		contentField, contentExists := gactusFile.Type().FieldByName("Content")

		if nameExists && contentExists && nameField.Type.Kind() == reflect.String && contentField.Type.Kind() == reflect.Slice {
			file, err := fileHeader.Open()
			if err != nil {
				return false, err
			}
			data, err := ioutil.ReadAll(file)
			if err != nil {
				return false, err
			}
			gactusFile.FieldByName("Name").SetString(fileHeader.Filename)
			gactusFile.FieldByName("Content").SetBytes(data)
			v.Set(gactusFile.Addr())
			return true, nil
		} else {
			return false, errors.New("GactusFile type does not have name and content field")
		}
	}
	return false, nil

}

func mapBasicDataType(v reflect.Value, value string) (found bool, err error) {
	kind := v.Kind()
	found = true
	switch kind {
	case reflect.Bool:
		data, err := strconv.ParseBool(value)
		if value != "" && err != nil {
			return false, err
		}
		v.SetBool(data)
	case reflect.Float32:
		data, err := strconv.ParseFloat(value, 32)
		if value != "" && err != nil {
			return false, err
		}
		v.SetFloat(data)
	case reflect.Float64:
		data, err := strconv.ParseFloat(value, 64)
		if value != "" && err != nil {
			return false, err
		}
		v.SetFloat(data)
	case reflect.Int:
		data, err := strconv.Atoi(value)
		if value != "" && err != nil {
			return false, err
		}
		v.SetInt(int64(data))
	case reflect.Int8:
		data, err := strconv.ParseInt(value, 10, 8)
		if value != "" && err != nil {
			return false, err
		}
		v.SetInt(data)
	case reflect.Int16:
		data, err := strconv.ParseInt(value, 10, 16)
		if value != "" && err != nil {
			return false, err
		}
		v.SetInt(data)
	case reflect.Int32:
		data, err := strconv.ParseInt(value, 10, 32)
		if value != "" && err != nil {
			return false, err
		}
		v.SetInt(data)
	case reflect.Int64:
		data, err := strconv.ParseInt(value, 10, 64)
		if value != "" && err != nil {
			return false, err
		}
		v.SetInt(data)
	case reflect.String:
		v.SetString(value)
	case reflect.Uint:
		data, err := strconv.ParseUint(value, 10, 32)
		if value != "" && err != nil {
			return false, err
		}
		v.SetUint(data)
	case reflect.Uint8:
		data, err := strconv.ParseUint(value, 10, 8)
		if value != "" && err != nil {
			return false, err
		}
		v.SetUint(data)
	case reflect.Uint16:
		data, err := strconv.ParseUint(value, 10, 16)
		if value != "" && err != nil {
			return false, err
		}
		v.SetUint(data)
	case reflect.Uint32:
		data, err := strconv.ParseUint(value, 10, 32)
		if value != "" && err != nil {
			return false, err
		}
		v.SetUint(data)
	case reflect.Uint64:
		data, err := strconv.ParseUint(value, 10, 64)
		if value != "" && err != nil {
			return false, err
		}
		v.SetUint(data)
	default:
		found = false
	}
	return found, nil
}
