package client

import (
	"errors"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"

	"github.com/cirocosta/sample-rpc-go/core"
)

// Client contains the configuration options for
// a RPC client that communicates with a RPC server
// over the network.
//
// Its parameters should match the server, for instance,
// if the server is offered via HTTP, it should have
// the property UseHttp set to true.
type Client struct {
	Port    uint
	UseHttp bool
	UseJson bool
	client  *rpc.Client
}

func (c *Client) Init() (err error) {
	if c.Port == 0 {
		err = errors.New("client: port must be specified")
		return
	}

	addr := "127.0.0.1:" + strconv.Itoa(int(c.Port))

	if c.UseHttp {
		c.client, err = rpc.DialHTTP("tcp", addr)
	} else if c.UseJson {
		c.client, err = jsonrpc.Dial("tcp", addr)
	} else {
		c.client, err = rpc.Dial("tcp", addr)
	}
	if err != nil {
		return
	}

	return
}

func (c *Client) Close() (err error) {
	if c.client != nil {
		err = c.client.Close()
		return
	}

	return
}

// TODO add context
func (c *Client) Execute(name string) (msg string, err error) {
	var (
		request  = &core.Request{Name: name}
		response = new(core.Response)
	)

	err = c.client.Call(core.HandlerName, request, response)
	if err != nil {
		return
	}

	msg = response.Message
	return

}
