package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SigningKey = "supersignkey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(SigningKey))
}

func ValidateToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid Token Signing method")
		}

		return []byte(SigningKey), nil
	})

	if err != nil {
		return errors.New("token Parsing Error")
	}

	IsTokenValid := parsedToken.Valid

	if !IsTokenValid {
		return errors.New("invalid token")
	}

	return nil
}
