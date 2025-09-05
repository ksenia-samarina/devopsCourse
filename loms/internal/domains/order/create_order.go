package order

import (
	"context"
	"devopsCourse/internal/domains/order/model"
	"fmt"

	"github.com/samber/lo"
)

func (d *Domain) CreateOrder(ctx context.Context, userID int64, items []model.Item) (int64, error) {
	var orderId int64

	err := d.transactionManager.RunWithLock(ctx, model.SoftLockLvl, func(ctxTX context.Context) error {
		var err error

		orderId, err = d.storage.CreateOrder(ctxTX, userID)
		if err != nil {
			return fmt.Errorf("error create new order: %v", err)
		}

		skus := lo.Map(items, func(val model.Item, _ int) uint32 { return val.Sku })

		availableStocks, err := d.storage.CntStocks(ctxTX, skus)
		if err != nil {
			err = d.storage.UpdateOrderStatus(ctxTX, orderId, model.FailedStatusOrder)
			if err != nil {
				return fmt.Errorf("error update order status to failed: %v",
					fmt.Errorf("error get stocks: %v", err))
			}

			return nil
		}

		skuStocks := lo.Reduce(availableStocks, func(agg map[uint32][]model.Stock, val model.Stock, _ int) map[uint32][]model.Stock {
			if data, ok := agg[val.Sku]; ok {
				agg[val.Sku] = append(data, val)
				return agg
			}

			agg[val.Sku] = []model.Stock{val}
			return agg

		}, make(map[uint32][]model.Stock, len(items)))

		reserveStocks, err := getReserveStocks(items, skuStocks)
		if err != nil {
			err = d.storage.UpdateOrderStatus(ctxTX, orderId, model.FailedStatusOrder)
			if err != nil {
				return fmt.Errorf("error update order status to failed: %v",
					fmt.Errorf("error get reserve stocks: %v", err))
			}

			return nil
		}

		err = d.storage.ReserveStocks(ctxTX, orderId, reserveStocks)
		if err != nil {
			err = d.storage.UpdateOrderStatus(ctxTX, orderId, model.FailedStatusOrder)
			if err != nil {
				return fmt.Errorf("error update order status to failed: %v",
					fmt.Errorf("error reserve stocks: %v", err))
			}

			return nil
		}

		err = d.storage.UpdateOrderStatus(ctxTX, orderId, model.AwaitingPaymentStatusOrder)
		if err != nil {
			return fmt.Errorf("error update order status: %v", err)
		}

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("error lock: %v", err)
	}

	return orderId, nil
}

func getReserveStocks(items []model.Item, skuStocks map[uint32][]model.Stock) ([]model.Stock, error) {
	reserveStocks := make([]model.Stock, 0, len(items))
	for _, item := range items {
		if data, ok := skuStocks[item.Sku]; ok {
			total := uint64(item.Count)

			for _, stock := range data {
				if stock.Count == 0 {
					continue
				}

				if total == 0 {
					break
				}

				var reserveCount uint64
				if total > stock.Count {
					total -= stock.Count
					reserveCount = stock.Count
				} else {
					reserveCount = total
					total = 0
				}
				reserveStocks = append(reserveStocks, model.Stock{
					Sku:         stock.Sku,
					WarehouseID: stock.WarehouseID,
					Count:       reserveCount,
				})
			}

			if total != 0 {
				return nil, fmt.Errorf("count reserve is large than in stocks, sku: %v", item.Sku)
			}
		} else {
			return nil, fmt.Errorf("stocks have not sku %d", item.Sku)
		}
	}

	return reserveStocks, nil
}
