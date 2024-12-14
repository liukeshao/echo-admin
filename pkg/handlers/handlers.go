package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/liukeshao/echo-admin/pkg/context"
)

func currentUserId(c echo.Context) uint64 {
	return c.Get(context.UserKey).(uint64)
}
