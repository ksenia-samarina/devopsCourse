package grpc_validator

import (
	"devopsCourse/internal/mistakes"
	lomsv1 "devopsCourse/pkg/loms_v1"
	"errors"
	"math"
)

func ValidateItem(value *lomsv1.Item) error {
	if value == nil {
		return errors.New("empty item")
	}

	if value.GetSku() == 0 {
		return mistakes.InvalidSku
	}

	if value.GetCount() == 0 || value.GetCount() > math.MaxUint16 {
		return errors.New("invalid count")
	}

	return nil
}
