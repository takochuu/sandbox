package main

import (
	"time"

	"github.com/acidlemon/go-dumper"
)

func main() {
	newServer("localhost", timeout(30), maxConnection(1))
}

type Server struct {
	timeout       time.Duration
	maxConnection int
}

func newServer(addr string, config ...func(*Server) error) *Server {
	s := &Server{}
	for _, v := range config {
		v(s)
	}
	dump.Dump(s)
	return s
}

func (s *Server) setTimeout(t time.Duration) error {
	s.timeout = t
	return nil
}

type Option func(*Server) error

func timeout(t int) Option {
	return func(s *Server) error {
		return s.setTimeout(time.Duration(t) * time.Second)
	}
}
func maxConnection(c int) Option {
	return func(s *Server) error {
		s.maxConnection = c
		return nil
	}
}
