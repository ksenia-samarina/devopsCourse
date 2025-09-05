package handlers

import (
	"devopsCourse/internal/domains/order"
	lomsv1 "devopsCourse/pkg/loms_v1"
)

type LomsServer struct {
	lomsv1.UnimplementedLomsV1Server

	OrderDomain order.Contract
}

func New(orderDomain order.Contract) *LomsServer {
	return &LomsServer{
		UnimplementedLomsV1Server: lomsv1.UnimplementedLomsV1Server{},
		OrderDomain:               orderDomain,
	}
}
