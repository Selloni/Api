package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	session "testgRPC/invoicer"
	"time"
)

func timingInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, // функция
	opts ...grpc.CallOption,
) error {
	start := time.Now() // обурнули все в тайминг
	err := invoker(ctx, method, req, reply, cc, opts...)
	// выводим всю инфрмацию
	log.Printf(`-- 
		call %v
		req %#v
		reply %#v
		time %v
		err %v`, method, req, reply, time.Since(start), err)
	return err
}

type tokenAuth struct { // реализовывает интерфейс
	Token string
}

// авторизация по токену
// эти данные будут добавляться к каждому запросу
func (t *tokenAuth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"access-token": t.Token,
	}, nil
}

// нужно ли для способа авторизации шифрование
func (t *tokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	grpConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithUnaryInterceptor(timingInterceptor),     // перехватчик запросов (midleweare)
		grpc.WithPerRPCCredentials(&tokenAuth{"100500"}), // передачие мето данных, по мимо резальтатов удаленных функций
		grpc.WithInsecure(),                              // не используем шифрование
	)

	if err != nil {
		log.Fatalln("cant connect to grpc")
	}
	defer grpConn.Close()

	sessionManager := session.NewAuthCheckerClient(grpConn)

	ctx := context.Background()
	// методанные передаються через контекст
	md := metadata.Pairs( // принимает четное количтево строк ключ-значение
		"api-req-id", "123",
		"subsystem", "cli",
	)
	ctx = metadata.NewOutgoingContext(ctx, md) // используется withvalue

	// ------

	var header, trailer metadata.MD // --даные которые придут в начале запроса и в конеце

	// создаем сессию
	sessId, err := sessionManager.Create(ctx,
		&session.Session{
			Login:     "Kot",
			Useragent: "chrome",
		},
		grpc.Header(&header), // передаем только системные данные
		grpc.Trailer(&trailer),
	)
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
