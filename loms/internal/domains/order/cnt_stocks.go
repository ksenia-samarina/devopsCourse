package order

import (
	"context"
	"devopsCourse/internal/domains/order/model"
	"fmt"

	"github.com/samber/lo"
)

func (d *Domain) CntStocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	stocks, err := d.storage.CntStocks(ctx, []uint32{sku})
	if err != nil {
		return nil, fmt.Errorf("error get stocks from db: %v", err)
	}

	stocksFiltered := lo.Filter(stocks, func(val model.Stock, _ int) bool {
		if val.Count == 0 {
			return false
		}

		return true
	})

	return stocksFiltered, nil
}
