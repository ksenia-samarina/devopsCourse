package grpc_validator

import (
	"devopsCourse/internal/mistakes"
	lomsv1 "devopsCourse/pkg/loms_v1"
)

func ValidateStocksIn(value *lomsv1.StocksIn) error {
	if value == nil {
		return mistakes.EmptyRequest
	}

	if value.GetSku() == 0 {
		return mistakes.InvalidSku
	}

	return nil
}
