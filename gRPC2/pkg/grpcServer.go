package pkg

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"math/rand"
	"strconv"
	"sync"
	session "testgRPC/invoicer"
)

type SessionManager struct {
	mu                                     sync.RWMutex                // блокировка
	sessions                               map[string]*session.Session // хранилище
	session.UnimplementedAuthCheckerServer                             // необходимо для реализации интерфейса
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		mu:       sync.RWMutex{},
		sessions: map[string]*session.Session{},
	}
}

func (s SessionManager) Create(ctx context.Context, in *session.Session) (*session.SessionId, error) {
	fmt.Println("call Create", in)
	id := strconv.Itoa(rand.Intn(98) + 1)
	fmt.Printf("create ID - %s\n", id)
	idS := &session.SessionId{ID: id}
	fmt.Println()
	s.mu.Lock()
	s.sessions[id] = in
	s.mu.Unlock()
	return idS, nil

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
	if _, ok := s.sessions[in.ID]; ok == true {
		delete(s.sessions, in.ID)
		return &session.Nothing{Dummy: true}, nil
	}
	return &session.Nothing{Dummy: false}, nil
}
