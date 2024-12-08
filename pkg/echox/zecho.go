package echox

import (
	z "github.com/Oudwins/zog"
	"github.com/labstack/echo/v4"
)

func Bind(c echo.Context, schema z.Schema, i any) error {
	var m map[string]any
	err := c.Bind(&m)
	if err != nil {
		return Error500InternalServerError("bind error", err)
	}
	errMap := z.Struct(schema).Parse(m, i)
	if errMap != nil {
		return Format(errMap)
	}
	return nil
}
