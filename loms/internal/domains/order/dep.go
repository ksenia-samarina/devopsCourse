package order

import (
	"context"
	"devopsCourse/internal/domains/order/model"
)

type Storage interface {
	CntStocks(ctx context.Context, skus []uint32) ([]model.Stock, error)
	CreateOrder(ctx context.Context, userID int64) (int64, error)
	UpdateOrderStatus(ctx context.Context, orderID int64, status model.StatusOrder) error
	UpdateStocksReservationStatus(ctx context.Context, orderID int64, status model.StatusStockReservation) error
	ListOrder(ctx context.Context, orderID int64) (model.ListOrder, error)
	ReserveStocks(ctx context.Context, orderId int64, reserveStocks []model.Stock) error
}

type TransactionManager interface {
	RunWithLock(ctx context.Context, lockLvl model.LockLvl, fx func(ctxTX context.Context) error) error
}
