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
type defaultClient struct {
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

func (c *defaultClient) fillConnPool(getConn bool) (conn *connection, err error) {
	c.poolLock.Lock()
	defer c.poolLock.Unlock()
	if c.poolSize == c.maxConns {
		return nil, errors.New("connection pool is full")
	}
	c.poolSize++
	tcpConn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return nil, err
	}
	conn = &connection{
		tcpConn:    tcpConn,
		lastActive: time.Now(),
	}
	if getConn {
		return conn, nil
	}
	c.connPool <- conn
	return nil, nil
}

func (c *defaultClient) poolManager() {
	for {
		if c.poolSize == 0 {
			return
		}
		conn := <-c.connPool
		if time.Since(conn.lastActive) > c.idleConnTimeout {
			_, _ = c.drainConnPool(conn, false)
		} else {
			c.connPool <- conn
		}
		time.Sleep(c.clearPeriod)
	}
}

// NewClient is used to create Client
func NewClient(addr string, minConns, maxConns int, idleConnTimeout, waitConnTimeout, clearPeriod time.Duration) (client Client, err error) {
	c := &defaultClient{
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
	for i := 0; i < c.minConns; i++ {
		if _, err = c.fillConnPool(false); err != nil {
			c.Close()
			return client, err
		}
	}
	go c.poolManager()
	return c, nil
}

// Send is used to send and get TCP data via TCP connection
func (c *defaultClient) Send(input []byte) (output []byte, err error) {
	conn := &connection{}
	select {
	case conn = <-c.connPool:
	case <-time.After(c.waitConnTimeout):
		conn, err = c.fillConnPool(true)
		if err != nil {
			return nil, err
		}
	}
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

func (c *defaultClient) drainConnPool(conn *connection, forceMode bool) (empty bool, err error) {
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
	c.poolSize--
	err = conn.tcpConn.Close()
	if err != nil {
		return false, err
	}
	empty = c.poolSize == 0
	return empty, nil
}

// Close is use to close all connections in pool
func (c *defaultClient) Close() (err error) {
	for empty := false; !empty; {
		empty, err = c.drainConnPool(nil, true)
		if err != nil {
			return err
		}
	}
	return nil
}
