package main

import (
	proto "gRPC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func main() {
	s := grpc.NewServer()
	srv := &proto.GRPCServer{} // реализованный интерфейс
	proto.RegisterAdderServer(s, srv)
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
