package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
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
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		sc := bufio.NewScanner(os.Stdin)
		fmt.Println("Введите текст, мы переведем его на английский")
		for sc.Scan() {
			txt := sc.Text()
			if txt == "" {
				client.CloseSend()
				return
			}
			client.Send(&session.Word{
				Word: txt,
			})
		}
	}(wg)

	// получаем данные
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			outWord, err := client.Recv()
			if err == io.EOF {
				fmt.Println("stream close")
				return
			}
			if err != nil {
				log.Fatal("fatal error:", err)
				return
			}
			fmt.Println("с_с):", outWord.Word)
		}
	}(wg)
	wg.Wait()
}
