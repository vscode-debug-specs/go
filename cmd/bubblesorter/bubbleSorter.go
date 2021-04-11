package main

import (
	"flag"
	"log"
	"net"

	"github.com/vscode-debug-specs/go/gateway/grpc"
)

func main() {
	var host string
	flag.StringVar(&host, "host", "0.0.0.0:8080", "host:port")
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	sv := grpc.New()
	err = sv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
