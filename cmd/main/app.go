package main

import (
	"RestApi/interal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("localHost:8080")
	router := httprouter.New()
	handler := user.NewHandler()
	handler.Register(router)
	run(router)
}
func run(router *httprouter.Router) {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}

}
