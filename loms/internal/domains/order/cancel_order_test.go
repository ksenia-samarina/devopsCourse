package order

import (
	"context"
	"devopsCourse/internal/domains/order/model"
	repomock "devopsCourse/internal/repository/postgres/mocks"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCancelOrder(t *testing.T) {
	t.Parallel()

	const orderID = int64(456)
	ctx := context.Background()

	tcs := []struct {
		name               string
		repoMock           func(mc *minimock.Controller) Storage
		transactionManager func() TransactionManager
		wantErr            string
	}{
		{
			name: "success cancel order",
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.UpdateOrderStatusMock.Expect(ctx, orderID, model.CancelledStatusOrder).Return(nil)
				mock.UpdateStocksReservationStatusMock.Expect(ctx, orderID, model.CanceledStatusStockReservation).Return(nil)

				return mock
			},
			transactionManager: func() TransactionManager {
				return &TransactionManagerImplMock{}
			},
			wantErr: "",
		},
		{
			name: "error update order status",
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.UpdateOrderStatusMock.Expect(ctx, orderID, model.CancelledStatusOrder).Return(errors.New("db fail"))

				return mock
			},
			transactionManager: func() TransactionManager {
				return &TransactionManagerImplMock{}
			},
			wantErr: "error lock: error update order status",
		},
		{
			name: "error update stocks reservation status",
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.UpdateOrderStatusMock.Expect(ctx, orderID, model.CancelledStatusOrder).Return(nil)
				mock.UpdateStocksReservationStatusMock.Expect(ctx, orderID, model.CanceledStatusStockReservation).Return(errors.New("oops"))

				return mock
			},
			transactionManager: func() TransactionManager {
				return &TransactionManagerImplMock{}
			},
			wantErr: "error lock: error update stocks reservation status: oops",
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
				transactionManager: tc.transactionManager(),
			}

			err := domain.CancelOrder(ctx, orderID)

			if tc.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantErr)
			}
		})
	}
}
