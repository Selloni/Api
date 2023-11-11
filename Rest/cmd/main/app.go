package main

import (
	author "RestApi/Rest/interal/book/db"
	"RestApi/Rest/interal/config"
	user2 "RestApi/Rest/interal/user"
	postrge "RestApi/Rest/pkg/client/postgresql"
	"RestApi/Rest/pkg/logging"
	"context"
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
	router := httprouter.New()
	cfg := config.GetConfig()

	client, err := postrge.NewClient(context.TODO(), 3, cfg.Storage)
	repository := author.NewRepository(client, logger)

	if err != nil {
		logger.Fatal(err)
	}

	//cfgMongo := cfg.MongoDB
	//mongoClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port,
	//	cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	//if err != nil {
	//	return
	//}
	//storage := db.NewStorage(mongoClient, cfgMongo.Collection, logger)
	//user1 := user2.User{
	//	Id:           "",
	//	Email:        "vip.petrusev@mail.ru",
	//	Username:     "lol",
	//	PasswordHash: "1234",
	//}
	//user1Id, err := storage.Create(context.Background(), user1)
	//logger.Info(user1Id)
	//
	//router := httprouter.New()
	//logger.Info("register user handler")
	handler := user2.NewHandler(logger)
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
