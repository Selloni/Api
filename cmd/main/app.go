package main

import (
	"RestApi/interal/user"
	"RestApi/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("start localhost:8080")

	router := httprouter.New()
	logger.Info("register user handler")
	handler := user.NewHandler()
	handler.Register(router)
	run(router)
}

func run(router *httprouter.Router) {
	logger := logging.GetLogger()
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		logger.Fatal(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := server.Serve(listen); err != nil {
		logger.Fatal(err)
	}

}
