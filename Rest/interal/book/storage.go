package book

import "context"

type Repository interface {
	Create(ctx context.Context, author *Author) error
	FindOne(ctx context.Context, id string) (Author, error)
	FindAll(ctx context.Context) (u []Author, err error)
}
