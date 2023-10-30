package adder

import (
	"context"
	proto "gRPC/proto"
)

type GRPCServer struct {
}

func (s *GRPCServer) Add(ctx context.Context, r *proto.AddRequest) (*proto.AddResponse, error) {
	return &proto.AddResponse{Result: r.GetX() + r.GetY()}, nil

}
