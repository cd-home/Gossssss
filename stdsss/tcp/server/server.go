package server

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type ServerFace interface {
	Start()
	Stop()
	Serve()
}

type Server struct {
	Name       string
	TCPVersion string
	IP         string
	Port       uint
}

func NewServer(name string) ServerFace {
	return &Server{
		Name:       name,
		TCPVersion: "tcp",
		IP:         "127.0.0.1",
		Port:       8080,
	}
}

func (s *Server) Start() {
	fmt.Printf("[start]--Listen IP: %s, Port: %d", s.IP, s.Port)
	addr, err := net.ResolveTCPAddr(s.TCPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP(s.TCPVersion, addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		// Connection
		var id uint = 0
		dealConn := NewConnection(conn, id, Callback)
		id++
		go dealConn.Start()
	}
}

func (s *Server) Stop() {
	// TODO
}

func (s *Server) Serve() {
	go s.Start()
	select {}
}

// client temp callback，later can design it self
func Callback(coon *net.TCPConn, data []byte, i int) error {
	if _, err := coon.Write(data); err != nil {
		log.Println(err)
		return errors.New("CallBack error")
	}
	return nil
}
