package pkg

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
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
	log.Println("call Create", in)

	// паралельное ip, все проверяется через runtime (пользоваться осторожно)
	header := metadata.Pairs("header-key", "42") // можно даставть данные из мапы
	grpc.SendHeader(ctx, header)                 // передаю методанные

	trailer := metadata.Pairs("trailer=key", "3.14")
	grpc.SetTrailer(ctx, trailer)

	id := strconv.Itoa(rand.Intn(98) + 1)
	idS := &session.SessionId{ID: id}
	s.mu.Lock()
	s.sessions[id] = in
	s.mu.Unlock()
	log.Printf("create ID - %s\n", idS.ID)
	return idS, nil
}
func (s SessionManager) Check(ctx context.Context, in *session.SessionId) (*session.Session, error) {
	log.Println("call Check", in)
	s.mu.RLock()
	if sess, ok := s.sessions[in.ID]; ok {
		return sess, nil
	}
	return nil, grpc.Errorf(codes.NotFound, "session not found")
}
func (s SessionManager) Delete(ctx context.Context, in *session.SessionId) (*session.Nothing, error) {
	log.Println("call Delete", in)
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.sessions[in.ID]; ok == true {
		delete(s.sessions, in.ID)
		return &session.Nothing{Dummy: true}, nil
	}
	return &session.Nothing{Dummy: false}, nil
}
