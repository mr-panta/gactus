package service

import (
	"net"

	"github.com/mr-panta/gactus/pkg/tcp"
)

// Handler [TOWRITE]
type Handler struct {
	coreClient *tcp.Client
}

// NewHandler [TOWRITE]
func NewHandler(coreAddr string, coreConnPoolSize int) *Handler {
	return &Handler{
		coreClient: tcp.NewClient(coreAddr, coreConnPoolSize), // TODO: pool size
	}
}

// ServeTCP is used to implement tcp.Handler
// and provides TCP connection.
func (h *Handler) ServeTCP(conn net.Conn) {}
