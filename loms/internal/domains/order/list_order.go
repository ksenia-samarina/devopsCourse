package order

import (
	"context"
	"devopsCourse/internal/domains/order/model"
	"fmt"
)

func (d *Domain) ListOrder(ctx context.Context, orderID int64) (model.ListOrder, error) {
	listOrder, err := d.storage.ListOrder(ctx, orderID)
	if err != nil {
		return model.ListOrder{}, fmt.Errorf("error get list order from db: %v", err)
	}

	if !listOrder.Status.IsValid() {
		return model.ListOrder{}, fmt.Errorf("invalid order status: %d, %s",
			listOrder.Status.Number(), listOrder.Status)
	}

	return listOrder, nil
}
