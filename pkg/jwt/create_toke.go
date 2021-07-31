package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

const EXPIRES_AFTER = 15000

func CreateToken(userId string) (string, error) {
	claims := ClaimsDTO{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + EXPIRES_AFTER,
			Issuer:    "banking-system",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("zdrthmfcv"))
	if err != nil {
		log.Error("Failed to parse token. Error: ", err)
		return "", err
	}
	return signedToken, nil
}
