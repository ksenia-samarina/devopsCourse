package transactor

import "github.com/jackc/pgx/v4/pgxpool"

type TransactionManager struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *TransactionManager {
	return &TransactionManager{
		pool: pool,
	}
}
