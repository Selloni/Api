package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/tap"
	"log"
	"net"
	session "testgRPC/invoicer"
	"testgRPC/pkg"
	"time"
)

func authInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	md, _ := metadata.FromIncomingContext(ctx) // методанные из запроса
	// в дальнешем можнореализовать аторизацию по токену

	reply, err := handler(ctx, req) // функция которая проделает реальную работу

	fmt.Printf(`--
		after incoming call=%v
		req=%#v
		reply=%#v
		time=%v
		md=%v
		err=%v
`, info.FullMethod, req, reply, time.Since(start), md, err)
	return reply, err
}

// ничего не распаршено, для проверки райтлимитов , посмортеь всю статистику до момента парсинга
// не стоит выполнять блокирующие операции, он выполняеться в одном потоке
func rateLimiter(ctx context.Context, info *tap.Info) (context.Context, error) {
	log.Printf("--\ncheck ratelim for %s\n", info.FullMethodName)
	return ctx, nil
}

// ---

func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("cannot crate listener: %v", err)
	}
	MyServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor), // перехватчик одиночных запросов
		grpc.InTapHandle(rateLimiter),          // выполяеть при приходе запроса на подключение(в самом начале),
	) // стандартный сервер
	service := pkg.NewSessionManager()                   // структура реализующая интерфейс
	session.RegisterAuthCheckerServer(MyServer, service) // сгенеренный код
	fmt.Println("starting server at:8081")
	MyServer.Serve(listen)
}
