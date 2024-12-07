package zogx

import "github.com/labstack/echo/v4"

func Request(c echo.Context) (map[string]any, error) {
	var m map[string]any
	err := c.Bind(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
