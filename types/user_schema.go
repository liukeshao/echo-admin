package types

import z "github.com/Oudwins/zog"

var CurrentUserSchema = z.Schema{
	"name": z.String().Required(),
}
