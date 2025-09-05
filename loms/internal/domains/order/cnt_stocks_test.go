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

func TestCntStocks(t *testing.T) {
	t.Parallel()

	const sku uint32 = 1
	ctx := context.Background()

	tcs := []struct {
		name     string
		repoMock func(mc *minimock.Controller) Storage
		want     []model.Stock
		wantErr  string
	}{
		{
			name: "success filter zero stocks",
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.CntStocksMock.Expect(ctx, []uint32{sku}).Return([]model.Stock{
					{Sku: sku, WarehouseID: 10, Count: 0},
					{Sku: sku, WarehouseID: 20, Count: 5},
				}, nil)

				return mock
			},
			want: []model.Stock{
				{Sku: sku, WarehouseID: 20, Count: 5},
			},
			wantErr: "",
		},
		{
			name: "error get stocks from db",
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.CntStocksMock.Expect(ctx, []uint32{sku}).Return(nil, errors.New("db fail"))

				return mock
			},
			want:    nil,
			wantErr: "error get stocks from db: db fail",
		},
		{
			name: "all stocks zero",
			repoMock: func(mc *minimock.Controller) Storage {
				mock := repomock.NewContractMock(mc)

				mock.CntStocksMock.Expect(ctx, []uint32{sku}).Return([]model.Stock{
					{Sku: sku, WarehouseID: 10, Count: 0},
					{Sku: sku, WarehouseID: 20, Count: 0},
				}, nil)

				return mock
			},
			want:    []model.Stock{}, // все отфильтровались
			wantErr: "",
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

			got, err := domain.CntStocks(ctx, sku)

			if tc.wantErr == "" {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			} else {
				require.EqualError(t, err, tc.wantErr)
			}
		})
	}
}
