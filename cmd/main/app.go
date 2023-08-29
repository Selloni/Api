package main

import (
	"RestApi/interal/config"
	"RestApi/interal/user"
	"RestApi/pkg/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("start localhost:8080")

	cfg := config.GetConfig()

	router := httprouter.New()
	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)
	run(router, cfg)
}

func run(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	var (
		ListenErr error
		appDir    string
		listener  net.Listener
		err       error
	)

	if cfg.Listen.Type == "sock" {
		appDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socker path %s", socketPath)
		logger.Info("create unix socket")
		listener, ListenErr = net.Listen("unix", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, ListenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}
	if ListenErr != nil {
		logger.Fatal(ListenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := server.Serve(listener); err != nil {
		logger.Fatal(err)
	}
	logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

}
