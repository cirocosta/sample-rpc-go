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

var (
	port        = flag.Uint("port", 1337, "port to listen or connect to for rpc calls")
	isServer    = flag.Bool("server", false, "activates server mode")
	json        = flag.Bool("json", false, "whether it should use json-rpc")
	serverSleep = flag.Duration("server.sleep", 0, "time for the server to sleep on requests")
	http        = flag.Bool("http", false, "whether it should use HTTP")
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

// must panics in the case of error.
func must(err error) {
	if err == nil {
		return
	}

	log.Panicln(err)
}

// runServer sets up the server with the
// flags as they were parsed and then initiates
// the server listening.
func runServer() {
	server := &Server{
		UseHttp: *http,
		UseJson: *json,
		Sleep:   *serverSleep,
		Port:    *port,
	}
	defer server.Close()

	go func() {
		handleSignals()
		server.Close()
		os.Exit(0)
	}()

	must(server.Start())
	return
}

// runClient sets up the client with the
// flags as they were parsed and then initiates
// the client execution.
func runClient() {
	client := &Client{
		UseHttp: *http,
		UseJson: *json,
		Port:    *port,
	}
	defer client.Close()

	must(client.Init())

	response, err := client.Execute("ciro")
	must(err)

	log.Println(response)
}

// main execution - validates flags and constructs the internal
// runtime configuration based on the flags supplied.
func main() {
	flag.Parse()

	if *isServer {
		log.Println("starting server")
		log.Printf("will listen on port %d\n", *port)

		runServer()
		return
	}

	log.Println("starting client")
	log.Printf("will connect to port %d\n", *port)

	runClient()
	return
}
