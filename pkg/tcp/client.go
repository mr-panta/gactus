package tcp

import (
	"io/ioutil"
	"net"
)

// Client contains TCP connection pool and provide
// APIs for communicating via the connections
type Client struct {
	addr     string
	connPool chan net.Conn
}

// NewClient is used to create Client
func NewClient(addr string, poolSize int) (client *Client) {
	client = &Client{addr: addr}
	client.connPool = make(chan net.Conn, poolSize)
	for i := 0; i < poolSize; i++ {
		conn, _ := client.newConn() // TODO: handle error
		client.connPool <- conn
	}
	return
}

func (c *Client) newConn() (conn net.Conn, err error) {
	return net.Dial("tcp", c.addr)
}

// Send is used to send and get TCP data via TCP connection
func (c *Client) Send(input []byte) (output []byte, err error) {
	conn := <-c.connPool
	defer func() {
		if err != nil {
			conn, _ = c.newConn()
		}
		c.connPool <- conn
	}()
	_, err = conn.Write(input)
	if err != nil {
		return nil, err
	}
	output, err = ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}
	return output, nil
}
