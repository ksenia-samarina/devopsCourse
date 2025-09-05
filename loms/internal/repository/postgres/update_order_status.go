package postgres

import (
	"context"
	"devopsCourse/internal/domains/order/model"
	"devopsCourse/internal/mistakes"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (s *Storage) UpdateOrderStatus(ctx context.Context, orderID int64, status model.StatusOrder) error {
	engine := s.QueryEngineProvider.GetQueryEngine(ctx)

	statusStorage, err := castOrderStatusToStorage(status)
	if err != nil {
		return fmt.Errorf("convert error: %v", err)
	}

	rawQuery, args, err := psql.Update(OrdersTable).
		Set(fieldOrdersStatus, statusStorage).
		Where(squirrel.Eq{fieldOrdersOrderId: orderID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("error buid query: %v", err)
	}

	result, err := engine.Exec(ctx, rawQuery, args...)
	if err != nil {
		return fmt.Errorf("error exec query: %v", err)
	}

	if result.RowsAffected() != 1 {
		return mistakes.NoAffectedRows
	}

	return nil
}
