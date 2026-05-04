package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtSigningSecret = "secret"

func JWTSigningKey() []byte {
	return []byte(jwtSigningSecret)
}

func NewJWTToken(subject string) (string, error) {
	claims := jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSigningKey())
}
