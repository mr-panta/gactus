package core

import (
	"errors"
	"io"
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
