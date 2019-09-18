package tcp

import (
	"context"
	"net"

	"github.com/mr-panta/gactus/pkg/logger"
)

// Handler is used to handle TCP connection.
type Handler interface {
	ServeTCP(conn net.Conn)
}

// ListenAndServe is used to listen to TCP connection.
func ListenAndServe(addr string, handler Handler) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		logger.Debugf(context.Background(), "new TCP connection is created")
		if err != nil {
			return err
		}
		go handler.ServeTCP(conn)
	}
}
