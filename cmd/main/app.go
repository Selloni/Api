package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	name := param.ByName("name")
	w.Write([]byte(fmt.Sprintf("hello %s", name)))
}

func main() {
	fmt.Printf("localHost:8080")
	router := httprouter.New()
	router.GET("/:name", IndexHandler)
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server.Serve(listen)
}
