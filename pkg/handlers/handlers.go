package handlers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

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
