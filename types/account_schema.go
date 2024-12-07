package types

import z "github.com/Oudwins/zog"

var LoginInputSchema = z.Struct(
	z.Schema{
		"email":    z.String().Email().Required(),
		"password": z.String().Required(),
	},
)

var RegisterInputSchema = z.Struct(
	z.Schema{
		"email":    z.String().Email().Required(),
		"password": z.String().Required(),
	},
)
