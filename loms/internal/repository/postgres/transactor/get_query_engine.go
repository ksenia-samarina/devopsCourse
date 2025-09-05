package transactor

import (
	"context"
	"devopsCourse/internal/repository/postgres"
)

func (tm *TransactionManager) GetQueryEngine(ctx context.Context) postgres.QueryEngine {
	tx, ok := ctx.Value(key).(postgres.QueryEngine)
	if ok && tx != nil {
		return tx
	}

	return tm.pool
}
