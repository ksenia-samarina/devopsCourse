package postgres

import (
	"context"
	"devopsCourse/internal/domains/order/model"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/samber/lo"
)

func (s *Storage) CntStocks(ctx context.Context, skus []uint32) ([]model.Stock, error) {
	engine := s.QueryEngineProvider.GetQueryEngine(ctx)

	aggQuery, aggArgs, err := psql.Select(
		fieldStocksReservationSku,
		fieldStocksWarehouseId,
		fmt.Sprintf("SUM(%s) as count", fieldStocksReservationCount),
	).From(StocksReservationTable).Where(squirrel.And{
		squirrel.Eq{fieldStocksReservationStatus: HoldStatusStockReservation},
		squirrel.Eq{fieldStocksReservationSku: skus},
	}).GroupBy(
		fieldStocksReservationSku,
		fieldStocksReservationWarehouseId,
	).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error buid query: %v", err)
	}

	placeholder := strings.Join(lo.Map(skus, func(val uint32, index int) string {
		return fmt.Sprintf("$%d", index+2)
	}), ",")

	rawQuery, _, err := psql.Select(
		fmt.Sprintf("%s.%s", StocksTable, fieldStocksSku),
		fmt.Sprintf("%s.%s", StocksTable, fieldStocksWarehouseId),
		fmt.Sprintf("COALESCE(%s.%s - t.count, %s.%s) as count", StocksTable, fieldStocksCount, StocksTable, fieldStocksCount),
	).From(StocksTable).LeftJoin(
		fmt.Sprintf(
			"(%s) as t on t.%s = %s.%s and t.%s = %s.%s",
			aggQuery,
			fieldStocksReservationWarehouseId,
			StocksTable,
			fieldStocksWarehouseId,
			fieldStocksReservationSku,
			StocksTable,
			fieldStocksSku),
	).Where(fmt.Sprintf("%s.%s in (%s)", StocksTable, fieldStocksSku, placeholder)).
		ToSql()

	var stocks []Stock
	err = pgxscan.Select(ctx, engine, &stocks, rawQuery, aggArgs...)
	if err != nil {
		return nil, fmt.Errorf("error select items: %v", err)
	}

	return lo.Map(stocks, func(val Stock, _ int) model.Stock {
		return model.Stock{
			Sku:         val.Sku,
			WarehouseID: val.WarehouseID,
			Count:       val.Count,
		}
	}), nil
}
