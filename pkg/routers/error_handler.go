package routers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/liukeshao/echo-admin/pkg/log"
	"net/http"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}
	logger := log.Ctx(c)
	logger.Error(err.Error())
	err = c.JSON(
		code,
		map[string]interface{}{
			"message": err.Error(),
		},
	)
	if err != nil {
		log.Ctx(c).Error("failed handle error",
			"error", err,
		)
	}
}
