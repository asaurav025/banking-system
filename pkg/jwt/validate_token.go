package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

func Validate(jwtFromHeader string) (string, error) {
	token, err := jwt.ParseWithClaims(
		jwtFromHeader,
		&ClaimsDTO{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("zdrthmfcv"), nil
		},
	)
	if err != nil {
		log.Error("Failed to parse with claim. Error: ", err)
		return "", err
	}
	claims, ok := token.Claims.(*ClaimsDTO)
	if !ok {
		log.Error("Failed to parse token. Error: ", err)
		return "", errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		log.Error("Token expired Error: ", err)
		return "", errors.New("jwt is expired")
	}

	return claims.UserId, nil
}
