package client

import (
	"errors"
	"net/rpc"
	"strconv"

	"github.com/cirocosta/sample-rpc/core"
)

type Client struct {
	Port uint
}

func (c *Client) Execute(name string) (msg string, err error) {
	var (
		client   *rpc.Client
		request  = &core.Request{Name: name}
		response = new(core.Response)
	)

	if c.Port == 0 {
		err = errors.New("client: port must be specified")
		return
	}

	addr := "127.0.0.1:" + strconv.Itoa(int(c.Port))

	client, err = rpc.Dial("tcp", addr)
	if err != nil {
		return
	}

	err = client.Call("Handler.SayHello", request, response)
	if err != nil {
		return
	}

	msg = response.Message
	return

}
