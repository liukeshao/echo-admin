package handlers

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Fail is a helper to fail a request by returning a 500 error and logging the error
func Fail(err error, log string) error {
	// The error handler will handle logging
	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", log, err))
}

func currentUserId(c echo.Context) (uint64, error) {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return 0, errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return 0, errors.New("failed to cast claims as jwt.MapClaims")
	}
	return uint64(claims["id"].(float64)), nil
}
