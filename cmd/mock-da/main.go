package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	proxy "github.com/rollkit/go-da/proxy/jsonrpc"
	goDATest "github.com/rollkit/go-da/test"
)

const (
	// MockDAPort is the port used for the mock DA gRPC server
	MockDAPort = 7980
)

var listenAll = flag.Bool("listen-all", false, "Listen on all network interfaces (0.0.0.0) instead of just localhost")

func main() {
	var (
		host string
		port string
	)
	flag.Parse()
        ip := "localhost"
        if *listenAll {
            ip = "0.0.0.0"
        }
	address := fmt.Sprintf("grpc://%s:%d", ip, MockDAPort)

	addr, _ := url.Parse(address)
	flag.StringVar(&port, "port", addr.Port(), "listening port")
	flag.StringVar(&host, "host", addr.Hostname(), "listening address")
	flag.Parse()

	srv := proxy.NewServer(host, port, goDATest.NewDummyDA())
	log.Printf("Listening on: %s:%s", host, port)
	if err := srv.Start(context.Background()); err != nil {
		log.Fatal("error while serving:", err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT)
	<-interrupt
	fmt.Println("\nCtrl+C pressed. Exiting...")
	os.Exit(0)
}
