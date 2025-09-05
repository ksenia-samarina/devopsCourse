package postgres

import (
	"context"
	order "devopsCourse/internal/domains/order/model"
	"fmt"

	"github.com/samber/lo"
)

func (s *Storage) ListOrder(ctx context.Context, orderID int64) (order.ListOrder, error) {
	orderDomain, err := s.GetOrder(ctx, orderID)
	if err != nil {
		return order.ListOrder{}, fmt.Errorf("error get order: %v", err)
	}

	items, err := s.GetStocksReservation(ctx, orderID)
	if err != nil {
		return order.ListOrder{}, fmt.Errorf("error get reservation stocks: %v", err)
	}

	return order.ListOrder{
		UserID: orderDomain.UserId,
		Status: orderDomain.Status,
		Items: lo.Map(items, func(val StocksReservation, _ int) order.Item {
			return order.Item{
				Sku:   uint32(val.Sku),
				Count: uint16(val.Count),
			}
		}),
	}, nil
}
