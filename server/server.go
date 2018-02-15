package server

import (
	"errors"
	"net"
	"net/rpc"
	"strconv"

	"github.com/cirocosta/sample-rpc/core"
)

type Server struct {
	Port     uint
	listener net.Listener
}

func (s *Server) Stop() (err error) {
	if s.listener != nil {
		err = s.listener.Close()
	}

	return
}

func (s *Server) Start() (err error) {
	if s.Port <= 0 {
		err = errors.New("port must be specified")
		return
	}

	rpc.Register(new(core.Handler))

	s.listener, err = net.Listen("tcp", ":"+strconv.Itoa(int(s.Port)))
	if err != nil {
		return
	}

	rpc.Accept(s.listener)
	return
}
