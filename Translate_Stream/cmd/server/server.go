package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	session "steam/api/proto"
	"steam/pkg/server"
)

func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("cannot crate listener: %v", err)
	}
	tr := server.TrServer{}                             // моя структура которая реализует интрефейс
	MyServer := grpc.NewServer()                        // стандартный сервер
	session.RegisterTransliterationServer(MyServer, tr) // сгенерированный код, для регистрации
	fmt.Println("starting server at:8081")
	MyServer.Serve(listen)
}
