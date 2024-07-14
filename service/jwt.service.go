package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTPayload struct {
	jwt.StandardClaims
	SessionID string
}

var (
	JWT_LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
	JWT_SIGNING_METHOD            = jwt.SigningMethodHS384
	JWT_SIGNATURE_KEY             = []byte(os.Getenv("JWT_SIGNATURE_KEY"))
)

func (payload JWTPayload) CreateJWTToken() (token string, err error, expiredDate int64) {
	var currentDate = time.Now()
	payload.IssuedAt = currentDate.Unix()
	payload.ExpiresAt = currentDate.Add(JWT_LOGIN_EXPIRATION_DURATION).Unix()
	expiredDate = payload.ExpiresAt
	jwtClaims := jwt.NewWithClaims(JWT_SIGNING_METHOD, payload)
	if token, err = jwtClaims.SignedString(JWT_SIGNATURE_KEY); err != nil {
		return
	}
	return
}

func ValidateJWTToken(tokenPayload string) (result jwt.MapClaims, err error) {
	tokenResult, err := jwt.Parse(tokenPayload, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return
	}

	result, ok := tokenResult.Claims.(jwt.MapClaims)
	if !ok || !tokenResult.Valid {
		err = errors.New("failed parse data")
		return
	}

	return
}
