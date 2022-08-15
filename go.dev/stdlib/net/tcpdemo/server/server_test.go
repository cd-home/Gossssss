package server

import (
	"net"
	"testing"
)

func TestTcp(t *testing.T) {
	conn, _ := net.Dial("tcpdemo", "localhost")
	conn.Close()
}
