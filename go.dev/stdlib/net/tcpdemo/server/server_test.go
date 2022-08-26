package server

import (
	"net"
	"testing"
)

func TestTcp(t *testing.T) {
	conn, _ := net.Dial("tcp", "localhost")
	conn.Close()
}
