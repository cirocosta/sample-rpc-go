package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	. "github.com/cirocosta/sample-rpc-go/client"
	. "github.com/cirocosta/sample-rpc-go/server"
)

func must(err error) {
	if err == nil {
		return
	}

	log.Panicln(err)
}

var (
	port     = flag.Uint("port", 1337, "port to listen or connect to for rpc calls")
	isServer = flag.Bool("server", false, "activates server mode")
	json     = flag.Bool("json", false, "whether it should use json-rpc")
	http     = flag.Bool("http", false, "whether it should use HTTP")
)

// handleSignals is a blocking function that waits for termination/interrupt
// signals.
//
// Running it in the background (non-main goroutine) has the effect of keeping
// track of the desire of termination of the current execution and then responding
// accordingly.
//
// In this example we gracefully  close the server listener in the case
// of the server - in the case of the client, breaks the request by cancelling the
// context.
func handleSignals() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Println("signal received")
}

func main() {
	flag.Parse()

	if *http {
		log.Println("will use http")
	}

	if *json {
		log.Println("will use json")
	}

	if *isServer {
		log.Println("starting server")
		log.Printf("will listen on port %d\n", *port)

		server := &Server{
			UseHttp: *http,
			UseJson: *json,
			Port:    *port,
		}

		go func() {
			handleSignals()
			server.Stop()
			os.Exit(0)
		}()

		must(server.Start())
		return
	}

	log.Println("starting client")
	log.Printf("will connect to port %d\n", *port)

	client := &Client{
		UseHttp: *http,
		UseJson: *json,
		Port:    *port,
	}

	response, err := client.Execute("ciro")
	must(err)

	log.Println(response)
}
