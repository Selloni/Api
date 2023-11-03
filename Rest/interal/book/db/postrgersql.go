package db

import (
	"RestApi/Rest/interal/book"
	"RestApi/Rest/pkg/client/postgresql"
	"RestApi/Rest/pkg/logging"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string { // что в логах поулчать запрос в комфортном виде
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, author *book.Author) error {
	q := `insert into author (name) values ($1) returning id`
	//q := `insert into author (name) values ($1) return id`
	// scan - записываем в поле возвращемое значение
	r.logger.Trace(fmt.Sprintf("SQL Qery :%s", formatQuery(q)))
	row := r.client.QueryRow(ctx, q, author.Name)
	if err := row.Scan(author.Id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %s, Defaiult %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}
func (r *repository) FindOne(ctx context.Context, id string) (book.Author, error) {
	var ath book.Author
	q := `select id, name from public.author where id = $1`
	r.logger.Trace(fmt.Sprintf("SQL Qery :%s", formatQuery(q)))
	row := r.client.QueryRow(ctx, q, id)
	err := row.Scan(&ath)
	if err != nil {
		return book.Author{}, err
	}

	if err = row.Err(); err != nil {
		return book.Author{}, err
	}
	return ath, nil
}
func (r *repository) FindAll(ctx context.Context) (u []book.Author, err error) {
	q := `select id, name from public.author`
	row, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	author := make([]book.Author, 0)
	for row.Next() {
		var ath book.Author
		err := row.Scan(&ath.Id, &ath.Name)
		if err != nil {
			return nil, err
		}
		author = append(author, ath)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	return author, nil
}

//func (r *repository) Update(ctx context.Context, author book.Author) error {
//
//}
//func (r *repository) Delete(ctx context.Context, id string) error {
//
//}

func NewRepository(client postgresql.Client, logger *logging.Logger) *repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
