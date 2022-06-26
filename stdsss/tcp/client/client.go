package client

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Client struct {
	Addr string
	Port uint
}

func NewClient(addr string, port uint) *Client {
	return &Client{
		Addr: addr,
		Port: port,
	}
}

func (c *Client) Send() {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", c.Addr, c.Port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = conn.Write([]byte("Hello cobar!"))
		if err != nil {
			log.Println(err)
			return
		}
		buffer := make([]byte, 1024)
		read, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Printf("server back message: %s", string(buffer[:read]))
		time.Sleep(5 * time.Second)
	}
}
