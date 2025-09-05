package grpc_validator

import (
	"devopsCourse/internal/mistakes"
	lomsv1 "devopsCourse/pkg/loms_v1"
	"errors"
	"fmt"
)

func ValidateCreateOrderIn(value *lomsv1.CreateOrderIn) error {
	if value == nil {
		return mistakes.EmptyRequest
	}

	if value.GetUser() <= 0 {
		return errors.New("invalid user")
	}

	items := value.GetItems()

	if len(items) == 0 {
		return errors.New("empty items")
	}

	var err error
	for ind, val := range items {
		err = ValidateItem(val)
		if err != nil {
			return fmt.Errorf("invalid items: item %d, %v", ind, err)
		}
	}

	return nil
}
