package interfaces

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IDBHandler interface {
	GetPool() *pgxpool.Pool
	AcquireConn(context.Context) (*pgxpool.Conn, error)
	StartTransaction(context.Context) (pgx.Tx, error)
	FinishTransaction(context.Context, pgx.Tx, error) error
}
