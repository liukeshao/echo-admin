package echox

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/liukeshao/echo-admin/pkg/log"
)

func ProblemHandler(err error, c echo.Context) {
	logger := log.Ctx(c)
	var se StatusError
	var he *echo.HTTPError
	if errors.As(err, &se) {
	} else if errors.As(err, &he) {
		se = NewError(he.Code, fmt.Sprintf("%+v", he.Message), he)
	} else {
		se = Error500InternalServerError("server internal error", err)
	}
	if se.GetStatus() >= http.StatusInternalServerError {
		problem := se.(*Problem)
		logger.Error(problem.Error(), "errorDetails", problem.Errors)
	}
	err = c.JSON(se.GetStatus(), se)
	if err != nil {
		log.Ctx(c).Error("failed handle error", "error", err)
	}
}
