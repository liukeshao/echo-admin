package echox

import (
	"net/http"
	"strings"

	z "github.com/Oudwins/zog"
)

func Format(errs z.ZogErrMap) *Problem {
	p := &Problem{
		Title:  http.StatusText(http.StatusBadRequest),
		Status: http.StatusBadRequest,
		Detail: "Validation failed",
	}
	sanitizeMap := z.Errors.SanitizeMap(errs)
	for filed, messages := range sanitizeMap {
		// https://zog.dev/errors
		// $root & $first are reserved keys for complex type validation, they are used for root level errors and for the first error found in a schema
		if filed == "$root" || filed == "$first" {
			continue
		}
		p.Add(&ErrorDetail{
			Message:  strings.Join(messages, ","),
			Location: filed,
			Value:    nil,
		})
	}
	return p
}
