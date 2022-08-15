package server

import (
	"fmt"
	"log"
	"net"
)

type ConnectionFace interface {
	Start()
	Stop()
	GetConnID() uint
	GetTCPConn() *net.TCPConn
	GetRemoteAddr() net.Addr
	Send([]byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error

type Connection struct {
	Conn   *net.TCPConn
	ConnID uint
	Close  bool
	F      HandleFunc
	Exit   chan bool
}

func NewConnection(conn *net.TCPConn, connID uint, callBack HandleFunc) *Connection {
	return &Connection{
		Conn:   conn,
		ConnID: connID,
		F:      callBack,
		Close:  false,
		Exit:   make(chan bool, 1),
	}
}

func (c *Connection) Reader() {
	log.Printf("conn [%d] is reading!", c.ConnID)
	defer fmt.Printf("conn [%s] reader is exit!", c.GetRemoteAddr().String())
	defer c.Stop()
	for {
		buffer := make([]byte, 1024)
		i, err := c.Conn.Read(buffer)
		if err != nil {
			log.Printf("conn [%d] recv error!", c.ConnID)
			continue
		}
		// CallBack
		if err := c.F(c.Conn, buffer, i); err != nil {
			log.Printf("conn [%d] handlefunc error!", c.ConnID)
			break
		}
	}
}

func (c *Connection) Start() {
	log.Printf("conn [%d] Start!", c.ConnID)
	// 启动读数据
	go c.Reader()
	// TODO 启动写
}

func (c *Connection) Stop() {
	log.Printf("conn [%d] Stop!", c.ConnID)
	if c.Close {
		return
	}
	c.Close = true
	_ = c.Conn.Close()
	close(c.Exit)
}

func (c *Connection) GetConnID() uint {
	return c.ConnID
}

func (c *Connection) GetTCPConn() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send([]byte) error {
	//TODO
	return nil
}
