package postgres

//go:generate sh -c "mkdir -p mocks && rm -rf mocks/*"
//go:generate minimock -g -i Contract -o ./mocks/ -s "_minimock.go"

import (
	"context"
	"devopsCourse/internal/domains/order/model"
)

type Contract interface {
	CntStocks(ctx context.Context, skus []uint32) ([]model.Stock, error)
	CreateOrder(ctx context.Context, userID int64) (int64, error)
	UpdateOrderStatus(ctx context.Context, orderID int64, status model.StatusOrder) error
	UpdateStocksReservationStatus(ctx context.Context, orderID int64, status model.StatusStockReservation) error
	ListOrder(ctx context.Context, orderID int64) (model.ListOrder, error)
	ReserveStocks(ctx context.Context, orderId int64, reserveStocks []model.Stock) error
}
