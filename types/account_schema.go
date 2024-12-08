package types

import z "github.com/Oudwins/zog"

var LoginInputSchema = z.Schema{
	"email":    z.String().Email().Required(),
	"password": z.String().Required(),
}

var RegisterInputSchema = z.Schema{
	"email":    z.String().Email().Required(),
	"password": z.String().Required(),
}
