package tcp

import (
	"errors"
	"io/ioutil"
	"net"
	"sync"
	"time"
)

// Client contains TCP connection pool and provide
// APIs for communicating via the connections
type Client struct {
	addr            string
	minConns        int
	maxConns        int
	idleConnTimeout time.Duration
	waitConnTimeout time.Duration
	clearPeriod     time.Duration
	poolSize        int
	poolLock        sync.Mutex
	connPool        chan *connection
}

type connection struct {
	tcpConn    net.Conn
	lastActive time.Time
}

func (c *Client) fillConnPool() (err error) {
	c.poolLock.Lock()
	defer c.poolLock.Unlock()
	if c.poolSize == c.maxConns {
		return errors.New("connection pool is full")
	}
	c.poolSize++
	tcpConn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.connPool <- &connection{
		tcpConn:    tcpConn,
		lastActive: time.Now(),
	}
	return nil
}

func (c *Client) poolManager() {
	for {
		if c.poolSize == 0 {
			return
		}
		conn := <-c.connPool
		if time.Since(conn.lastActive) > c.idleConnTimeout {
			_, _ = c.drainConnPool(conn, false)
		}
		time.Sleep(c.clearPeriod)
	}
}

// NewClient is used to create Client
func NewClient(addr string, minConns, maxConns int, idleConnTimeout, waitConnTimeout, clearPeriod time.Duration) (client *Client, err error) {
	client = &Client{
		addr:            addr,
		minConns:        minConns,
		maxConns:        maxConns,
		idleConnTimeout: idleConnTimeout,
		waitConnTimeout: waitConnTimeout,
		clearPeriod:     clearPeriod,
		poolSize:        0,
		poolLock:        sync.Mutex{},
		connPool:        make(chan *connection, maxConns),
	}
	for i := 0; i < client.minConns; i++ {
		if err = client.fillConnPool(); err != nil {
			client.Close()
			return client, err
		}
	}
	go client.poolManager()
	return client, nil
}

// Send is used to send and get TCP data via TCP connection
func (c *Client) Send(input []byte) (output []byte, err error) {
	conn := <-c.connPool
	conn.lastActive = time.Now()
	defer func() {
		c.connPool <- conn
	}()
	_, err = conn.tcpConn.Write(input)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(conn.tcpConn)
}

func (c *Client) drainConnPool(conn *connection, forceMode bool) (empty bool, err error) {
	c.poolLock.Lock()
	defer c.poolLock.Unlock()
	if c.poolSize == 0 {
		return true, errors.New("connection pool is empty")
	}
	if conn == nil {
		conn = <-c.connPool
	}
	defer func() {
		if err != nil {
			c.connPool <- conn
		}
	}()
	if c.poolSize == c.minConns && !forceMode {
		err = errors.New("pool size cannot be lower than minimum number of connections")
		return false, err
	}
	err = conn.tcpConn.Close()
	if err != nil {
		return false, err
	}
	empty = c.poolSize == 0
	return empty, nil
}

// Close is use to close all connections in pool
func (c *Client) Close() (err error) {
	for empty := false; !empty; {
		empty, err = c.drainConnPool(nil, true)
		if err != nil {
			return err
		}
	}
	return nil
}
