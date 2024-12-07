package services

import (
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

// GenerateToken generates a token for a user id using JWT which
// is set to expire based on the duration stored in configuration
func (s *AuthService) GenerateToken(userId uint64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  userId,
			"exp": time.Now().Add(s.config.App.EncryptionExpiration).Unix(),
		},
	)

	return token.SignedString([]byte(s.config.App.EncryptionKey))
}
