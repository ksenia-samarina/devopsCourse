package postgres

import (
	"context"
	order "devopsCourse/internal/domains/order/model"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (s *Storage) GetOrder(ctx context.Context, orderId int64) (order.Order, error) {
	engine := s.QueryEngineProvider.GetQueryEngine(ctx)

	rawQuery, args, err := psql.Select(ordersFields...).From(OrdersTable).
		Where(squirrel.Eq{fieldOrdersOrderId: orderId}).
		ToSql()
	if err != nil {
		return order.Order{}, fmt.Errorf("error buid query: %v", err)
	}

	var orderDB Order
	err = pgxscan.Get(ctx, engine, &orderDB, rawQuery, args...)
	if err != nil {
		return order.Order{}, fmt.Errorf("error select items: %v", err)
	}

	return order.Order{
		OrderId:   orderDB.OrderId,
		UserId:    orderDB.UserId,
		Status:    castOrderStatusToDomainOrder(orderDB.Status),
		CreatedAt: orderDB.CreatedAt,
	}, nil
}
