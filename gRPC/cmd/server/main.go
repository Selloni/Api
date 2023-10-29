package main

import (
	"gRPC/pkg/adder"
	"gRPC/pkg/api/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{} // реализованный интерфейс
	proto.RegisterAdderServer(s, srv)
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
