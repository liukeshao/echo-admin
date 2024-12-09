package echox

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/liukeshao/echo-admin/pkg/log"
)

func ProblemHandler(err error, c echo.Context) {
	var problem *Problem
	var he *echo.HTTPError
	if errors.As(err, &problem) {
	} else if errors.As(err, &he) {
		problem = NewError(he.Code, fmt.Sprintf("%+v", he.Message), he)
	} else {
		problem = Error500InternalServerError("server internal error", err)
	}
	log.Ctx(c).Error(problem.Error(), "errorDetails", problem.Errors)
	err = c.JSON(problem.GetStatus(), problem)
	if err != nil {
		log.Ctx(c).Error("failed handle error", "error", err)
	}
}
