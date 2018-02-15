package server

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/cirocosta/sample-rpc/core"
)

type Server struct {
	Port     uint
	UseHttp  bool
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

	if s.UseHttp {
		rpc.HandleHTTP()
		http.Serve(s.listener, nil)
	} else {
		rpc.Accept(s.listener)
	}

	return
}
