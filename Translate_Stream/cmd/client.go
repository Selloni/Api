package main

import (
	"google.golang.org/grpc"
	"log"
)

func main() {
	grpcConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc")
	}
	defer grpcConn.Close()
}
