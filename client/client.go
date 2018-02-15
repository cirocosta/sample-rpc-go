package client

import (
	"context"
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

// Init initializes the underlying RPC client that is
// responsible for taking a codec and writing the RPC
// details down to it.
//
// Here we're using basic `(json)rpc.Dial*` but we could
// instead make use of raw TCP connections and tune them
// accordingly if we'd take this to production as we could
// then specify various timeouts and options for the transport
// layer.
//
// Note.: the HTTP thing is just a very thin layer of HTTP
// that is sent via the TCP connection. We could ditch that
// of and replace by a `CONNECT` call followed by checking
// the HTTP response that we got back.
//
// Note.: we're not setting TLS here either but it's a very
// simple thing given that we can have total control over
// the underlying connection.
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

// Close gracefully terminates the underlying client.
func (c *Client) Close() (err error) {
	if c.client != nil {
		err = c.client.Close()
		return
	}

	return
}

// Executes the only handler that we want to execute (`core.Handler`).
//
// Given that `net/rpc` has been freezed and is not taking more
// features, it didn't get a `context` nativelly, which is sad.
//
// To have the benefits of `context` we can then wrap `client.Call()`
// with this `Execute` method and provide the proper deadline
// enforcement and the other cancellation features.
func (c *Client) Execute(ctx context.Context, name string) (msg string, err error) {
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
