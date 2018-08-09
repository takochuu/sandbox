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

func newServer(addr string, config ...func(*Server)) *Server {
	s := &Server{}
	for _, v := range config {
		v(s)
	}
	dump.Dump(s)
	return s
}

func (s *Server) setTimeout(t time.Duration) {
	s.timeout = t
}

type Option func(*Server)

func timeout(t int) Option {
	return func(s *Server) {
		s.setTimeout(time.Duration(t) * time.Second)
	}
}
func maxConnection(c int) Option {
	return func(s *Server) {
		s.maxConnection = c
	}
}
