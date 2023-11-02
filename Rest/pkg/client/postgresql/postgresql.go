package postgresql

import (
	"RestApi/interal/config"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int, config config.StorageConfig) (*pgxpool.Pool, error) {
	var (
		pool *pgxpool.Pool
		err  error
	)
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.Username, config.Password,
		config.Host, config.Port, config.Database)
	for maxAttempts > 0 { // для подчключения к бд
		ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
		defer cancelFunc()
		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return nil, fmt.Errorf("failed connect to DB")
		}
		time.Sleep(5 * time.Second)
		maxAttempts--
	}
	return pool, nil
}
