package postgres

import (
	"time"

	"github.com/Masterminds/squirrel"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

const (
	OrdersTable            = "orders"
	StocksTable            = "stocks"
	StocksReservationTable = "stocks_reservation"
)

var (
	fieldOrdersUserId  = "user_id"
	fieldOrdersStatus  = "status"
	fieldOrdersOrderId = "order_id"
	fieldCreatedAt     = "created_at"
)

var (
	ordersFields            = []string{fieldOrdersOrderId, fieldOrdersUserId, fieldOrdersStatus, fieldCreatedAt}
	stocksReservationFields = []string{
		fieldStocksReservationOrderId,
		fieldStocksReservationSku,
		fieldStocksReservationWarehouseId,
		fieldStocksReservationCount,
		fieldStocksReservationStatus,
	}
)

type Order struct {
	OrderId   int64       `db:"order_id"`
	UserId    int64       `db:"user_id"`
	Status    StatusOrder `db:"status"`
	CreatedAt time.Time   `db:"created_at"`
}

type StatusOrder uint8

const (
	NewStatusOrder             = StatusOrder(1) // new
	AwaitingPaymentStatusOrder = StatusOrder(2) // awaiting payment
	FailedStatusOrder          = StatusOrder(3) // failed
	PaidStatusOrder            = StatusOrder(4) // paid
	CancelledStatusOrder       = StatusOrder(5) // cancelled
)

var (
	fieldStocksReservationOrderId     = "order_id"
	fieldStocksReservationSku         = "sku"
	fieldStocksReservationWarehouseId = "warehouse_id"
	fieldStocksReservationCount       = "count"
	fieldStocksReservationStatus      = "status"
)

type StocksReservation struct {
	OrderId     int64                  `db:"order_id"`
	Sku         int64                  `db:"sku"`
	WarehouseId int64                  `db:"warehouse_id"`
	Count       int32                  `db:"count"`
	Status      StatusStockReservation `db:"status"`
}

type StatusStockReservation uint8

const (
	HoldStatusStockReservation     = StatusStockReservation(1) // hold
	PaidStatusStockReservation     = StatusStockReservation(2) // paid
	CanceledStatusStockReservation = StatusStockReservation(3) // cancelled
)

var (
	fieldStocksSku         = "sku"
	fieldStocksWarehouseId = "warehouse_id"
	fieldStocksCount       = "count"
)

type Stock struct {
	Sku         uint32 `db:"sku"`
	WarehouseID int64  `db:"warehouse_id"`
	Count       uint64 `db:"count"`
}
