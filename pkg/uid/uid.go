package uid

import (
	"github.com/oklog/ulid/v2"
)

// ULID https://github.com/ulid/spec
func ULID() string {
	return ulid.Make().String()
}
