package handlers

import (
	"context"
	grpcvalidator "devopsCourse/internal/validators/grpc_validator"
	lomsv1 "devopsCourse/pkg/loms_v1"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (ls *LomsServer) CancelOrder(ctx context.Context, in *lomsv1.CancelOrderIn) (*emptypb.Empty, error) {
	err := grpcvalidator.ValidateCancelOrderIn(in)
	if err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	err = ls.OrderDomain.CancelOrder(ctx, in.GetOrderID())
	if err != nil {
		log.Println("error cancel order:", err)
		return nil, fmt.Errorf("error cancel order: %v", err)
	}

	return &emptypb.Empty{}, nil
}
