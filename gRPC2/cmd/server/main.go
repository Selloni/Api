package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	session "testgRPC/invoicer"
	"testgRPC/pkg"
)

func main() {

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("cannot crate listener: %v", err)
	}
	MyServer := grpc.NewServer()
	service := pkg.NewSessionManager()
	session.RegisterAuthCheckerServer(MyServer, service)
	fmt.Println("starting server at:8081")
	MyServer.Serve(listen)

}
