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

func TestListOrder(t *testing.T) {
	t.Parallel()

	const orderID int64 = 123
	ctx := context.Background()

	tcs := []struct {
		name     string
		repoMock func(mc *minimock.Controller) Storage
		want     model.ListOrder
		wantErr  string
	}{
		{
			name: "success list order",
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.ListOrderMock.Expect(ctx, orderID).Return(model.ListOrder{
					UserID: 42,
					Status: model.AwaitingPaymentStatusOrder, // валидный статус
					Items: []model.Item{
						{Sku: 1, Count: 2},
					},
				}, nil)

				return mock
			},
			want: model.ListOrder{
				UserID: 42,
				Status: model.AwaitingPaymentStatusOrder,
				Items: []model.Item{
					{Sku: 1, Count: 2},
				},
			},
			wantErr: "",
		},
		{
			name: "error from db",
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.ListOrderMock.Expect(ctx, orderID).Return(model.ListOrder{}, errors.New("db fail"))

				return mock
			},
			want:    model.ListOrder{},
			wantErr: "error get list order from db: db fail",
		},
		{
			name: "invalid order status",
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				invalidStatus := model.StatusOrder(6)

				mock.ListOrderMock.Expect(ctx, orderID).Return(model.ListOrder{
					UserID: 42,
					Status: invalidStatus,
					Items:  nil,
				}, nil)

				return mock
			},
			want:    model.ListOrder{},
			wantErr: "invalid order status: 6, unknown status",
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mc := minimock.NewController(t)
			defer mc.Finish()

			domain := Domain{
				storage: tc.repoMock(mc),
			}

			got, err := domain.ListOrder(ctx, orderID)

			if tc.wantErr == "" {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			} else {
				require.EqualError(t, err, tc.wantErr)
			}
		})
	}
}
