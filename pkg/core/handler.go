package core

import (
	"net"
	"net/http"
)

type httpHandler struct{}

// ServeHTTP is used to implement http.Handler,
// get HTTP request and send back HTTP response
func (h *httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {}

type tcpHandler struct{}

// ServeTCP is used to implement tcp.Handler
// and provides TCP connection
func (h *tcpHandler) ServeTCP(conn net.Conn) {}
