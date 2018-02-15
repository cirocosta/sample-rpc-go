package core

import (
	"errors"
)

type Response struct {
	Message string
	Ok      bool
}

type Request struct {
	Name string
}

type Handler struct{}

func (h *Handler) SayHello(req Request, res *Response) (err error) {
	if req.Name == "" {
		err = errors.New("A name must be specified")
		return
	}

	res.Ok = true
	res.Message = "Hello " + req.Name

	return
}
