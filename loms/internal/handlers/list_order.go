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

func (ls *LomsServer) ListOrder(ctx context.Context, in *lomsv1.ListOrderIn) (*lomsv1.ListOrderOut, error) {
	err := grpcvalidator.ValidateListOrderIn(in)
	if err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	listOrder, err := ls.OrderDomain.ListOrder(ctx, in.GetOrderID())
	if err != nil {
		log.Println("error getting order list:", err)
		return nil, fmt.Errorf("error getting order list: %s", err.Error())
	}

	result := castToListOrderOut(listOrder)
	return result, nil
}

func castToListOrderOut(listOrder model.ListOrder) *lomsv1.ListOrderOut {
	items := lo.Map(listOrder.Items, func(val model.Item, _ int) *lomsv1.Item {
		return &lomsv1.Item{
			Sku:   val.Sku,
			Count: uint32(val.Count),
		}
	})

	var status lomsv1.OrderStatus
	switch listOrder.Status {
	case model.NewStatusOrder:
		status = lomsv1.OrderStatus_STATUS_NEW
	case model.AwaitingPaymentStatusOrder:
		status = lomsv1.OrderStatus_STATUS_AWAITING_PAYMENT
	case model.FailedStatusOrder:
		status = lomsv1.OrderStatus_STATUS_FAILED
	case model.PaidStatusOrder:
		status = lomsv1.OrderStatus_STATUS_PAYED
	case model.CancelledStatusOrder:
		status = lomsv1.OrderStatus_STATUS_CANCELLED
	default:
		status = lomsv1.OrderStatus_STATUS_UNKNOWN
	}

	return &lomsv1.ListOrderOut{
		User:   listOrder.UserID,
		Status: status,
		Items:  items,
	}
}
