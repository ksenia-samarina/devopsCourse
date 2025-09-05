package postgres

import (
	"context"
	order "devopsCourse/internal/domains/order/model"
	"devopsCourse/internal/mistakes"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (s *Storage) UpdateStocksReservationStatus(ctx context.Context, orderID int64, status order.StatusStockReservation) error {
	engine := s.QueryEngineProvider.GetQueryEngine(ctx)

	statusReservation, err := castStatusStockReservationToStorage(status)
	if err != nil {
		return fmt.Errorf("cast error: %v", err)
	}

	rawQuery, args, err := psql.Update(StocksReservationTable).
		Set(fieldStocksReservationStatus, statusReservation).
		Where(squirrel.Eq{fieldStocksReservationOrderId: orderID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("error buid query: %v", err)
	}

	result, err := engine.Exec(ctx, rawQuery, args...)
	if err != nil {
		return fmt.Errorf("error exec query: %v", err)
	}

	if result.RowsAffected() == 0 {
		return mistakes.NoAffectedRows
	}

	return nil
}
