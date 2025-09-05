package handlers

import (
	"context"
	"devopsCourse/internal/domains/order/model"
	grpcvalidator "devopsCourse/internal/validators/grpc_validator"
	lomsv1 "devopsCourse/pkg/loms_v1"
	"fmt"
	"log"

	"github.com/samber/lo"
)

func (ls *LomsServer) CntStocks(ctx context.Context, in *lomsv1.StocksIn) (*lomsv1.StocksOut, error) {
	err := grpcvalidator.ValidateStocksIn(in)
	if err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	stocks, err := ls.OrderDomain.CntStocks(ctx, in.GetSku())
	if err != nil {
		log.Println("error get stocks:", err)
		return nil, fmt.Errorf("error get stocks: %s", err.Error())
	}

	outStocks := lo.Map(stocks, func(val model.Stock, _ int) *lomsv1.StockItem {
		return &lomsv1.StockItem{
			WarehouseID: val.WarehouseID,
			Count:       val.Count,
		}
	})

	return &lomsv1.StocksOut{
		Stocks: outStocks,
	}, nil
}
