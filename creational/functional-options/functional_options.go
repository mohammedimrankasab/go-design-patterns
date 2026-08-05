package functionaloptions

import "time"

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
	Logging bool
}

type Option func(*Server)

func NewServer(options ...Option) *Server {
	s := &Server{
		Host:    "localhost",
		Port:    8080,
		Timeout: 30 * time.Second,
		Logging: false,
	}

	for _, option := range options {
		option(s)
	}

	return s
}

func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithLogging(logging bool) Option {
	return func(s *Server) {
		s.Logging = logging
	}
}
