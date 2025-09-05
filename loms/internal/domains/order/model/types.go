package model

import "time"

type ListOrder struct {
	UserID int64
	Status StatusOrder
	Items  []Item
}

type Order struct {
	OrderId   int64
	UserId    int64
	Status    StatusOrder
	CreatedAt time.Time
}

type Item struct {
	Sku   uint32
	Count uint16
}

type StatusOrder uint8

const (
	NewStatusOrder             = StatusOrder(1) // new
	AwaitingPaymentStatusOrder = StatusOrder(2) // awaiting payment
	FailedStatusOrder          = StatusOrder(3) // failed
	PaidStatusOrder            = StatusOrder(4) // paid
	CancelledStatusOrder       = StatusOrder(5) // cancelled
)

func (so StatusOrder) Number() uint8 {
	return uint8(so)
}

func (so StatusOrder) String() string {
	switch so {
	case NewStatusOrder:
		return "new"
	case AwaitingPaymentStatusOrder:
		return "awaiting payment"
	case FailedStatusOrder:
		return "failed"
	case PaidStatusOrder:
		return "paid"
	case CancelledStatusOrder:
		return "cancelled"
	default:
		return "unknown status"
	}
}

func (so StatusOrder) IsValid() bool {
	switch so {
	case NewStatusOrder, AwaitingPaymentStatusOrder, FailedStatusOrder, PaidStatusOrder, CancelledStatusOrder:
		return true
	}

	return false
}

type Stock struct {
	Sku         uint32
	WarehouseID int64
	Count       uint64
}

type LockLvl uint8

const (
	SoftLockLvl   = LockLvl(0)
	MiddleLockLvl = LockLvl(1)
	HardLockLvl   = LockLvl(2)
)

type StatusStockReservation uint8

const (
	HoldStatusStockReservation     = StatusStockReservation(1) // hold
	PaidStatusStockReservation     = StatusStockReservation(2) // paid
	CanceledStatusStockReservation = StatusStockReservation(3) // cancelled
)

func (ssr StatusStockReservation) Number() uint8 {
	return uint8(ssr)
}

func (ssr StatusStockReservation) String() string {
	switch ssr {
	case HoldStatusStockReservation:
		return "hold"
	case PaidStatusStockReservation:
		return "payed"
	case CanceledStatusStockReservation:
		return "cancelled"
	default:
		return "unknown status"
	}
}
