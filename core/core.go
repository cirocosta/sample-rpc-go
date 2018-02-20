// core implements shared functionality that both
// client and server can use.
//
// By exporting all the messages (Request and Response)
// it becomes very easy for the client to communicate
// back and forth with the server.
package core

import (
	"errors"
	"time"
)

type Response struct {
	Message string
	Ok      bool
}

type Request struct {
	Name string
}

// HandlerName provider the name of the only
// method that `core` exposes via the RPC
// interface.
//
// This could be replaced by the use of the reflect
// package (e.g, `reflect.ValueOf(func).Pointer()).Name()`).
const HandlerName = "Handler.Execute"

// Handler holds the methods to be exposed by the RPC
// server as well as properties that modify the methods'
// behavior.
type Handler struct {

	// Sleep adds a little sleep between to the
	// method execution to simulate a time-consuming
	// operation.
	Sleep time.Duration
}

// Execute is the exported method that a RPC client can
// make use of by calling the RPC server using `HandlerName`
// as the endpoint.
//
// It takes a Request and produces a Response if no error
// happens, possibly sleeping in between if a sleep is
// specified in Handler.
func (h *Handler) Execute(req Request, res *Response) (err error) {
	if req.Name == "" {
		err = errors.New("A name must be specified")
		return
	}

	if h.Sleep != 0 {
		time.Sleep(h.Sleep)
	}

	res.Ok = true
	res.Message = "Hello " + req.Name

	return
}
