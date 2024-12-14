package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/liukeshao/echo-admin/pkg/context"
	"github.com/liukeshao/echo-admin/pkg/services"
	"net/http"
)

// RequireAuthentication requires that the user must be authenticated in order to proceed
func RequireAuthentication(auth *services.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get(echo.HeaderAuthorization)
			if authorization == "" {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			claim, err := auth.ParseToken(authorization)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			c.Set(context.UserKey, claim.UserId)
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
