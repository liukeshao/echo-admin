package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/liukeshao/echo-admin/pkg/context"
)

func currentUserId(c echo.Context) string {
	return c.Get(context.UserKey).(string)
}
