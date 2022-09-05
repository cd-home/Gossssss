package snippetsss

import "testing"

type Server struct {
	Pprof bool
	Debug bool
}

type Option interface {
	apply(*Server)
}

type OptionFunc func(*Server)

func (f OptionFunc) apply(s *Server) {
	f(s)
}

// Default true
func AddPprof() Option {
	return WithPprof(true)
}

func WithPprof(enabled bool) Option {
	return OptionFunc(func(s *Server) {
		s.Pprof = enabled
	})
}

func WithDebug(enabled bool) Option {
	return OptionFunc(func(s *Server) {
		s.Debug = enabled
	})
}

func NewServer(opts ...Option) (*Server, error) {
	s := &Server{}
	s, _ = s.WithOptions(opts...)
	return s, nil
}

func (s *Server) clone() *Server {
	copy := *s
	return &copy
}

func (s *Server) WithOptions(opts ...Option) (*Server, error) {
	c := s.clone()
	for _, opt := range opts {
		opt.apply(c)
	}
	return c, nil
}

func TestOptionMode(t *testing.T) {
	server, _ := NewServer(AddPprof())
	t.Log(server.Pprof)

	server, _ = server.WithOptions(WithDebug(true))
	t.Log(server.Debug)
}
