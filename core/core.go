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

const HandlerName = "Handler.Execute"

type Handler struct {
	Sleep time.Duration
}

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
