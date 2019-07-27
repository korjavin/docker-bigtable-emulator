package main

import (
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/bigtable/bttest"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8086, "the port number to bind to on the local machine")
)

const (
	maxMsgSize = 256 * 1024 * 1024 // 256 MiB
)

func main() {
	grpc.EnableTracing = false
	flag.Parse()
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(maxMsgSize),
		grpc.MaxSendMsgSize(maxMsgSize),
	}
	srv, err := bttest.NewServer(fmt.Sprintf(":%d", *port), opts...)
	if err != nil {
		log.Fatalf("failed to start emulator: %v", err)
	}

	fmt.Printf("Cloud Bigtable emulator running on %s\n", srv.Addr)
	select {}
}
