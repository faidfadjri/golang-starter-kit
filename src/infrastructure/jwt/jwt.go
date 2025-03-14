package jwtutil

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID   uint64   `json:"user_id"`
	Fullname string `json:"fullname,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role"`
	Phone    *string `json:"phone,omitempty"`
	Address  *string `json:"address,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
	jwt.RegisteredClaims
}

func GenerateToken(user Claims, expiresIn time.Duration) (string, error) {
	user.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
