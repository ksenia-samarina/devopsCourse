package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (s *Storage) GetStocksReservation(ctx context.Context, orderId int64) ([]StocksReservation, error) {
	engine := s.QueryEngineProvider.GetQueryEngine(ctx)

	rawQuery, args, err := psql.
		Select(stocksReservationFields...).
		From(StocksReservationTable).
		Where(squirrel.Eq{fieldOrdersOrderId: orderId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error buid query: %v", err)
	}

	var items []StocksReservation
	err = pgxscan.Select(ctx, engine, &items, rawQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("error select items: %v", err)
	}

	return items, nil
}
