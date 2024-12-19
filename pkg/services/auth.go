package services

import (
	"errors"
	"fmt"
	"github.com/liukeshao/echo-admin/types"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/liukeshao/echo-admin/config"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	config *config.Config
}

func NewAuthService(config *config.Config) *AuthService {
	service := new(AuthService)
	service.config = config
	return service
}

// HashPassword returns a hash of a given password
func (s *AuthService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword check if a given password matches a given hash
func (s *AuthService) CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// GenToken generates a token for a user id using JWT which
// is set to expire based on the duration stored in configuration
func (s *AuthService) GenToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.MyCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.App.EncryptionExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			ID:        "",
			Audience:  []string{},
		},
	})

	// Sign and get the complete encoded token as a string using the key
	tokenString, err := token.SignedString([]byte(s.config.App.EncryptionKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) ParseToken(tokenString string) (*types.MyCustomClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	if strings.TrimSpace(tokenString) == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&types.MyCustomClaims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(s.config.App.EncryptionKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*types.MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
