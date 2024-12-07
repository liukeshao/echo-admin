package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// RequireAuthentication requires that the user must be authenticated in order to proceed
func RequireAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Get(echo.HeaderAuthorization)
			if authorization == nil {
				return echo.NewHTTPError(http.StatusForbidden)
			}
			return next(c)
		}
	}
}

// RequireNoAuthentication requires that the user not be authenticated in order to proceed
func RequireNoAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Get(echo.HeaderAuthorization)
			if authorization != nil {
				return echo.NewHTTPError(http.StatusForbidden)
			}
			return next(c)
		}
	}
}
