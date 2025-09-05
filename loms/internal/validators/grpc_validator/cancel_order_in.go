package grpc_validator

import (
	"devopsCourse/internal/mistakes"
	lomsv1 "devopsCourse/pkg/loms_v1"
)

func ValidateCancelOrderIn(value *lomsv1.CancelOrderIn) error {
	if value == nil {
		return mistakes.EmptyRequest
	}

	if value.GetOrderID() <= 0 {
		return mistakes.InvalidOrderID
	}

	return nil
}
