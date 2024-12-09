package types

import z "github.com/Oudwins/zog"

var OrgCreateInputSchema = z.Schema{
	"name":         z.String().Required(),
	"display_name": z.String().Optional(),
	"description":  z.String().Optional(),
	"logo":         z.String().Optional(),
	"status":       z.String().Optional(),
	"type":         z.String().Optional(),
}
