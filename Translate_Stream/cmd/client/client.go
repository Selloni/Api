package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	session "steam/api/proto"
	"sync"
)

func main() {
	grpcConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal("cant connect to grpc")
	}
	defer grpcConn.Close()

	transit := session.NewTransliterationClient(grpcConn)

	ctx := context.Background()
	client, err := transit.EnRu(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// отправлем данные
	func(wg *sync.WaitGroup) {
		defer wg.Done()
		ruWord := []string{"Привет", "удивительный", "МИР", "!"}
		for _, s := range ruWord {
			client.Send(&session.Word{
				Word: s,
			})
			//time.Sleep(1 * time.Second)
		}
		client.CloseSend()
	}(wg)

	// получаем данные
	func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			outWord, err := client.Recv()
			if err == io.EOF {
				fmt.Println("stream close")
				return
			}
			if err != nil {
				fmt.Println("fatal error:", err)
				return
			}
			fmt.Println(outWord)
		}
	}(wg)
	wg.Wait()
}
