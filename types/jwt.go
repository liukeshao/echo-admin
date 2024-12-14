package types

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	jwt.RegisteredClaims

	UserId uint64 `json:"user_id"`
}
