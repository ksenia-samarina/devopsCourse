package order

import (
	"context"
	"devopsCourse/internal/domains/order/model"
)

type Contract interface {
	CancelOrder(ctx context.Context, orderID int64) error
	CreateOrder(ctx context.Context, userID int64, items []model.Item) (int64, error)
	ListOrder(ctx context.Context, orderID int64) (model.ListOrder, error)
	CntStocks(ctx context.Context, sku uint32) ([]model.Stock, error)
}
