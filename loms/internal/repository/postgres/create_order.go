package postgres

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
)

func (s *Storage) CreateOrder(ctx context.Context, userID int64) (int64, error) {
	engine := s.QueryEngineProvider.GetQueryEngine(ctx)

	rawQuery, args, err := psql.Insert(OrdersTable).
		Columns(fieldOrdersUserId, fieldOrdersStatus).
		Values(userID, NewStatusOrder).
		Suffix("RETURNING order_id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("error buid query: %v", err)
	}

	var orderId int64
	err = pgxscan.Get(ctx, engine, &orderId, rawQuery, args...)
	if err != nil {
		return 0, err
	}

	return orderId, nil
}
