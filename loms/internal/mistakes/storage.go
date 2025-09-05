package mistakes

import "errors"

var (
	NoAffectedRows = errors.New("rows no affected")
)
