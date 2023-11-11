package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	session "testgRPC/invoicer"
)

func main() {
	grpConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(), // не используем шифрование
	)
	if err != nil {
		log.Fatalln("cant connect to grpc")
	}
	defer grpConn.Close()
	sessionManager := session.NewAuthCheckerClient(grpConn)
	ctx := context.Background()

	// создаем сессию
	sessId, err := sessionManager.Create(ctx,
		&session.Session{
			Login:     "Kot",
			Useragent: "chrome",
		})
	log.Println("sessId", sessId, err)

	// проверяем сессию
	sess, err := sessionManager.Check(ctx,
		&session.SessionId{ID: sessId.ID})
	log.Println("sess", sess, err)

	// удаляем сессию
	_, err = sessionManager.Delete(ctx,
		&session.SessionId{ID: sessId.ID})

	// проверяем  еще раз сессию
	sess, err = sessionManager.Check(ctx,
		&session.SessionId{ID: sessId.ID})
	log.Println("sess", sess, err)
}
