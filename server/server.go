// server Implements an RPC server that exposes
// a hello-world to RPC clients.
//
// The server is accessible either via HTTP, TCP
// or JSON-RPC.
package server

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"time"

	"github.com/cirocosta/sample-rpc-go/core"
)

// Server holds the configuration used to initiate
// an RPC server.
type Server struct {
	Port     uint
	UseHttp  bool
	UseJson  bool
	Sleep    time.Duration
	listener net.Listener
}

// Close gracefully terminates the server listener.
func (s *Server) Close() (err error) {
	if s.listener != nil {
		err = s.listener.Close()
	}

	return
}

// Starts initializes the RPC server by first verifying
// if all the necessary configuration has been set.
//
// It then publishes the receiver's methods (core.Handler)
// in the default RPC server. By doing so, the `Handler`
// public methods that satisfy the rpc interface become
// available to clients connecting to this server.
//
// With the receiver registered, it starts the server
// such that new connections can be accepted.
func (s *Server) Start() (err error) {
	if s.Port <= 0 {
		err = errors.New("port must be specified")
		return
	}

	rpc.Register(&core.Handler{
		Sleep: s.Sleep,
	})

	s.listener, err = net.Listen("tcp", ":"+strconv.Itoa(int(s.Port)))
	if err != nil {
		return
	}

	if s.UseHttp {
		rpc.HandleHTTP()
		http.Serve(s.listener, nil)
	} else if s.UseJson {
		var conn net.Conn

		for {
			conn, err = s.listener.Accept()
			if err != nil {
				return
			}

			jsonrpc.ServeConn(conn)
		}

	} else {
		rpc.Accept(s.listener)
	}

	return
}
