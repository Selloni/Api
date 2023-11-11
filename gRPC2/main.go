package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"math/rand"
	"net"
	"sync"
	session "testgRPC/invoicer"
)

type SessionManager struct {
	mu       sync.RWMutex
	sessions map[string]*session.Session
	session.UnimplementedAuthCheckerServer
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		mu:       sync.RWMutex{},
		sessions: map[string]*session.Session{},
	}
}

func (s SessionManager) Create(ctx context.Context, in *session.Session) (*session.SessionId, error) {
	fmt.Println("call Create", in)
	id := &session.SessionId{ID: string(rand.Intn(99-1) + 1)}
	s.mu.Lock()
	s.sessions[id.ID] = in
	s.mu.Unlock()
	return id, nil

}
func (s SessionManager) Check(ctx context.Context, in *session.SessionId) (*session.Session, error) {
	fmt.Println("call Check", in)
	s.mu.RLock()
	if sess, ok := s.sessions[in.ID]; ok {
		return sess, nil
	}
	return nil, grpc.Errorf(codes.NotFound, "session not found")
}
func (s SessionManager) Delete(ctx context.Context, in *session.SessionId) (*session.Nothing, error) {
	fmt.Println("call Delete", in)
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, in.ID)
	return &session.Nothing{Dummy: true}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot crate listener: %v", err)
	}
	MyServer := grpc.NewServer()
	service := &SessionManager{}
	session.RegisterAuthCheckerServer(MyServer, service)
	fmt.Println("starting server at:8081")
	MyServer.Serve(listen)

}