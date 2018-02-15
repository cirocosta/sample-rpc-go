
        DESCRIPTION

                sample-rpc-go provides an example of a client-server
                communication making use of `net/rpc` and `net/rpc/jsonrpc`.

                Both client and server are meant to be run using a single
                binary - to run a server, turn the `-server` flag on:

                        ./main -server

                There are three modes supported:

                        - default:      binary RPC (encoding/gob) via pure TCP
                        - http:         binary RPC (encoding/god) via HTTP
                        - json:         JSON-RPC (1.0) via TCP

                These modes are meant to be used together (server and client should
                be on the same mode) and can be activated with the respective flags.


        USAGE

                ./main --help
                Usage of ./main:
                  -http
                        whether it should use HTTP
                  -json
                        whether it should use json-rpc
                  -port uint
                        port to listen or connect to for rpc calls (default 1337)
                  -server
                        activates server mode


        HACK

                # Produce a binary called `main` in the repository root
                make build

                # Format
                make fmt

