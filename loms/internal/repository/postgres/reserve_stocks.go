package postgres

import (
	"context"
	order "devopsCourse/internal/domains/order/model"
	"devopsCourse/internal/mistakes"
	"fmt"
)

func (s *Storage) ReserveStocks(ctx context.Context, orderId int64, reserveStocks []order.Stock) error {
	engine := s.QueryEngineProvider.GetQueryEngine(ctx)

	insertQuery := psql.Insert(StocksReservationTable).
		Columns(stocksReservationFields...)

	for _, stock := range reserveStocks {
		insertQuery = insertQuery.Values(
			orderId,
			stock.Sku,
			stock.WarehouseID,
			stock.Count,
			HoldStatusStockReservation,
		)
	}

	rawQuery, args, err := insertQuery.ToSql()
	if err != nil {
		return fmt.Errorf("error buid query: %v", err)
	}

	result, err := engine.Exec(ctx, rawQuery, args...)
	if err != nil {
		return fmt.Errorf("error exec query: %v", err)
	}

	if result.RowsAffected() != int64(len(reserveStocks)) {
		return mistakes.NoAffectedRows
	}

	return nil
}
