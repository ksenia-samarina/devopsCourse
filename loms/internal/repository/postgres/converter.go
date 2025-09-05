package postgres

import (
	order "devopsCourse/internal/domains/order/model"
	"fmt"
)

func castOrderStatusToDomainOrder(status StatusOrder) order.StatusOrder {
	switch status {
	case NewStatusOrder:
		return order.NewStatusOrder
	case AwaitingPaymentStatusOrder:
		return order.AwaitingPaymentStatusOrder
	case FailedStatusOrder:
		return order.FailedStatusOrder
	case PaidStatusOrder:
		return order.PaidStatusOrder
	case CancelledStatusOrder:
		return order.CancelledStatusOrder
	default:
		return order.StatusOrder(0) // invalid value
	}
}

func castOrderStatusToStorage(status order.StatusOrder) (StatusOrder, error) {
	switch status {
	case order.NewStatusOrder:
		return NewStatusOrder, nil
	case order.AwaitingPaymentStatusOrder:
		return AwaitingPaymentStatusOrder, nil
	case order.FailedStatusOrder:
		return FailedStatusOrder, nil
	case order.PaidStatusOrder:
		return PaidStatusOrder, nil
	case order.CancelledStatusOrder:
		return CancelledStatusOrder, nil
	default:
		return StatusOrder(0), fmt.Errorf("invalid status order: %s %d", status.String(), status.Number())
	}
}

func castStatusStockReservationToStorage(status order.StatusStockReservation) (StatusStockReservation, error) {
	switch status {
	case order.HoldStatusStockReservation:
		return HoldStatusStockReservation, nil
	case order.PaidStatusStockReservation:
		return PaidStatusStockReservation, nil
	case order.CanceledStatusStockReservation:
		return CanceledStatusStockReservation, nil

	default:
		return StatusStockReservation(0), fmt.Errorf("invalid status stock reservation: %s %d", status.String(), status.Number())
	}
}
