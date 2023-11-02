package main

import (
	"context"
	proto "gRPC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	proto.UnimplementedAdderServer
}

func (s *GRPCServer) Add(ctx context.Context, r *proto.AddRequest) (*proto.AddResponse, error) {
	return &proto.AddResponse{Result: r.GetX() + r.GetY()}, nil
}

func main() {
	s := grpc.NewServer()
	srv := &GRPCServer{} // реализованный интерфейс
	proto.RegisterAdderServer(s, srv)
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
