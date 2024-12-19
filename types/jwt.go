package types

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	jwt.RegisteredClaims

	UserId string `json:"user_id"`
}
