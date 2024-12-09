package errors

import "github.com/joomcode/errorx"

var (
	RecordNotFound = errorx.CommonErrors.NewType("record not found", errorx.NotFound())
)
