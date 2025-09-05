package order

import (
	"context"
	"devopsCourse/internal/domains/order/model"
	repomock "devopsCourse/internal/repository/postgres/mocks"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

type TransactionManagerImplMock struct {
}

func (t *TransactionManagerImplMock) RunWithLock(ctx context.Context, lockLvl model.LockLvl, fx func(ctxTX context.Context) error) error {
	return fx(ctx)
}

func TestCreateOrder(t *testing.T) {
	t.Parallel()

	const userID = int64(123)
	const orderID = int64(456)

	var (
		items = []model.Item{
			{
				Sku:   uint32(1),
				Count: uint16(1),
			},
			{
				Sku:   uint32(2),
				Count: uint16(2),
			},
			{
				Sku:   uint32(3),
				Count: uint16(3),
			},
		}
		ctx = context.Background()
	)

	type args struct {
		ctx    context.Context
		userID int64
		items  []model.Item
	}

	tcs := []struct {
		name               string
		args               args
		repoMock           func(mc *minimock.Controller) Storage
		transactionManager func(mc *minimock.Controller) TransactionManager
		want               int64
		wantErr            string
	}{
		{
			name: "error create order",
			args: args{
				ctx:    ctx,
				userID: userID,
				items:  items,
			},
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.CreateOrderMock.Expect(ctx, userID).Return(0, errors.New("oops"))

				return mock
			},
			want:    0,
			wantErr: "error lock: error create new order: oops",
		},
		{
			name: "error get stocks",
			args: args{
				ctx:    ctx,
				userID: userID,
				items:  items,
			},
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.CreateOrderMock.Expect(ctx, userID).Return(orderID, nil)

				skus := lo.Map(items, func(val model.Item, _ int) uint32 { return val.Sku })
				mock.CntStocksMock.Expect(ctx, skus).Return(nil, errors.New("oops"))

				mock.UpdateOrderStatusMock.Expect(ctx, orderID, model.FailedStatusOrder).Return(errors.New("oops"))

				return mock
			},
			want:    0,
			wantErr: "error lock: error update order status to failed: error get stocks: oops",
		},
		{
			name: "success order creation",
			args: args{ctx, userID, items},
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.CreateOrderMock.Expect(ctx, userID).Return(orderID, nil)

				mock.CntStocksMock.Expect(ctx, []uint32{uint32(1), uint32(2), uint32(3)}).Return([]model.Stock{
					{Sku: uint32(1), WarehouseID: int64(10), Count: uint64(2)},
					{Sku: uint32(2), WarehouseID: int64(10), Count: uint64(2)},
					{Sku: uint32(3), WarehouseID: int64(10), Count: uint64(3)},
				}, nil)

				mock.ReserveStocksMock.Expect(ctx, orderID, []model.Stock{
					{Sku: uint32(1), WarehouseID: int64(10), Count: uint64(1)},
					{Sku: uint32(2), WarehouseID: int64(10), Count: uint64(2)},
					{Sku: uint32(3), WarehouseID: int64(10), Count: uint64(3)},
				}).Return(nil)

				mock.UpdateOrderStatusMock.Expect(ctx, orderID, model.AwaitingPaymentStatusOrder).Return(nil)

				return mock
			},
			want:    orderID,
			wantErr: "",
		},
		{
			name: "error getReserveStocks not enough stock",
			args: args{ctx, userID, items},
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.CreateOrderMock.Expect(ctx, userID).Return(orderID, nil)

				mock.CntStocksMock.Expect(ctx, []uint32{uint32(1), uint32(2), uint32(3)}).Return([]model.Stock{
					{Sku: 1, WarehouseID: 10, Count: 1},
					{Sku: 2, WarehouseID: 10, Count: 1},
					{Sku: 3, WarehouseID: 10, Count: 1},
				}, nil)

				mock.UpdateOrderStatusMock.Expect(ctx, orderID, model.FailedStatusOrder).Return(nil)

				return mock
			},
			want:    orderID,
			wantErr: "",
		},
		{
			name: "error reserve stocks",
			args: args{ctx, userID, items},
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.CreateOrderMock.Expect(ctx, userID).Return(orderID, nil)

				mock.CntStocksMock.Expect(ctx, []uint32{uint32(1), uint32(2), uint32(3)}).Return([]model.Stock{
					{Sku: 1, WarehouseID: 10, Count: 1},
					{Sku: 2, WarehouseID: 10, Count: 2},
					{Sku: 3, WarehouseID: 10, Count: 3},
				}, nil)

				mock.ReserveStocksMock.Expect(ctx, orderID, []model.Stock{
					{Sku: 1, WarehouseID: 10, Count: 1},
					{Sku: 2, WarehouseID: 10, Count: 2},
					{Sku: 3, WarehouseID: 10, Count: 3},
				}).Return(errors.New("reserve fail"))

				mock.UpdateOrderStatusMock.Expect(ctx, orderID, model.FailedStatusOrder).Return(nil)

				return mock
			},
			want:    orderID,
			wantErr: "",
		},
		{
			name: "error final update status",
			args: args{ctx, userID, items},
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.CreateOrderMock.Expect(ctx, userID).Return(orderID, nil)

				mock.CntStocksMock.Expect(ctx, []uint32{uint32(1), uint32(2), uint32(3)}).Return([]model.Stock{
					{Sku: 1, WarehouseID: 10, Count: 1},
					{Sku: 2, WarehouseID: 10, Count: 2},
					{Sku: 3, WarehouseID: 10, Count: 3},
				}, nil)

				mock.ReserveStocksMock.Expect(ctx, orderID, []model.Stock{
					{Sku: 1, WarehouseID: 10, Count: 1},
					{Sku: 2, WarehouseID: 10, Count: 2},
					{Sku: 3, WarehouseID: 10, Count: 3},
				}).Return(nil)

				mock.UpdateOrderStatusMock.Expect(ctx, orderID, model.AwaitingPaymentStatusOrder).Return(errors.New("db fail"))

				return mock
			},
			want:    0,
			wantErr: "error lock: error update order status: db fail",
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mc := minimock.NewController(t)
			defer mc.Finish()

			domain := Domain{
				storage:            tc.repoMock(mc),
				transactionManager: &TransactionManagerImplMock{},
			}

			got, gotErr := domain.CreateOrder(tc.args.ctx, tc.args.userID, tc.args.items)

			if tc.wantErr == "" {
				require.NoError(t, gotErr)
				require.Equal(t, tc.want, got)
			} else {
				require.EqualError(t, gotErr, tc.wantErr)
				require.Empty(t, got)
			}
		})
	}
}
