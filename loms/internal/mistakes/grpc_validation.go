package mistakes

import "errors"

var (
	EmptyRequest   = errors.New("empty request")
	InvalidOrderID = errors.New("invalid orderID")
	InvalidSku     = errors.New("invalid sku")
)
