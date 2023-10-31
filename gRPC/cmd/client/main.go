package main

import (
	"context"
	"fmt"
	"gRPC/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	con, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		return
	}
	fmt.Print("oooo")
	c := proto.NewAdderClient(con)
	res, err := c.Add(context.Background(), &proto.AddRequest{X: 2, Y: 17})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("ffff")
	log.Print(res.GetResult())
}
