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
	MyServer := grpc.NewServer()                         // стандартный сервер
	service := pkg.NewSessionManager()                   // структура реализующая интерфейс
	session.RegisterAuthCheckerServer(MyServer, service) // сгенеренный код
	fmt.Println("starting server at:8081")
	MyServer.Serve(listen)
}
