package infrastructure

import (
	"api/internal/config"
	"api/internal/interfaces"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type PostgresClient struct {
	Pool *pgxpool.Pool
}

func initPostgresClient(cfg *config.Config, ctx context.Context) (interfaces.IDBHandler, error) {
	dbPool, err := pgxpool.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if err != nil {
		return nil, errors.Wrap(err, "postgres init")
	}
	return &PostgresClient{Pool: dbPool}, nil
}

func (p *PostgresClient) GetPool() *pgxpool.Pool {
	return p.Pool
}

func (p *PostgresClient) AcquireConn(ctx context.Context) (*pgxpool.Conn, error) {
	return p.Pool.Acquire(ctx)
}

func (p *PostgresClient) StartTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Begin")
	}

	return tx, err
}

func (p *PostgresClient) FinishTransaction(ctx context.Context, tx pgx.Tx, err error) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return errors.Wrap(err, "Rollback")
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			return errors.Wrap(err, "failed to commit tx")
		}

		return nil
	}
}
