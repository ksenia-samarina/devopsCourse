package transactor

import (
	"context"
	"devopsCourse/internal/domains/order/model"

	"github.com/jackc/pgx/v4"
	"go.uber.org/multierr"
)

func (tm *TransactionManager) RunWithLock(ctx context.Context, lockLvl model.LockLvl, fx func(ctxTX context.Context) error) error {
	pgLockLvl := chooseLockLvl(lockLvl)

	tx, err := tm.pool.BeginTx(
		ctx,
		pgx.TxOptions{
			IsoLevel: pgLockLvl,
		})
	if err != nil {
		return err
	}

	if err := fx(context.WithValue(ctx, key, tx)); err != nil {
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	if err := tx.Commit(ctx); err != nil {
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	return nil
}

func chooseLockLvl(in model.LockLvl) pgx.TxIsoLevel {
	switch in {
	case model.SoftLockLvl:
		return pgx.ReadCommitted
	case model.MiddleLockLvl:
		return pgx.RepeatableRead
	case model.HardLockLvl:
		return pgx.Serializable
	}

	return pgx.ReadUncommitted
}
