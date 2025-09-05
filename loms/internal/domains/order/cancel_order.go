package order

import (
	"context"
	"devopsCourse/internal/domains/order/model"
	"fmt"
)

func (d *Domain) CancelOrder(ctx context.Context, orderID int64) error {
	err := d.transactionManager.RunWithLock(ctx, model.SoftLockLvl, func(ctxTX context.Context) error {
		err := d.storage.UpdateOrderStatus(ctxTX, orderID, model.CancelledStatusOrder)
		if err != nil {
			return fmt.Errorf("error update order status")
		}

		err = d.storage.UpdateStocksReservationStatus(ctxTX, orderID, model.CanceledStatusStockReservation)
		if err != nil {
			return fmt.Errorf("error update stocks reservation status: %v", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error lock: %v", err)
	}

	return nil
}
