package transactor

import (
	"context"
	"devopsCourse/internal/domains/order/model"
	"devopsCourse/internal/repository/postgres"
)

type Contract interface {
	RunWithLock(ctx context.Context, lockLvl model.LockLvl, fx func(ctxTX context.Context) error) error
	GetQueryEngine(ctx context.Context) postgres.QueryEngine
}
