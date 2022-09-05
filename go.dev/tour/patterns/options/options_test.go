package options

import "testing"

func TestOptionMode(t *testing.T) {
	server, _ := NewServer(AddPprof())
	t.Log(server.Pprof)

	server, _ = server.WithOptions(WithDebug(true))
	t.Log(server.Debug)
}
