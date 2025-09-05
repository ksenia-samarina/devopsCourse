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

func (ls *LomsServer) CreateOrder(ctx context.Context, in *lomsv1.CreateOrderIn) (*lomsv1.CreateOrderOut, error) {
	err := grpcvalidator.ValidateCreateOrderIn(in)
	if err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	items := lo.Map(in.GetItems(), func(item *lomsv1.Item, _ int) model.Item {
		return model.Item{
			Sku:   item.Sku,
			Count: uint16(item.GetCount()),
		}
	})

	orderID, err := ls.OrderDomain.CreateOrder(ctx, in.GetUser(), items)
	if err != nil {
		log.Println("error create order:", err)
		return nil, fmt.Errorf("error create order: %s", err.Error())
	}

	return &lomsv1.CreateOrderOut{
		OrderID: orderID,
	}, nil
}
