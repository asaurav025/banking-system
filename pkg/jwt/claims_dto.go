package jwt

import "github.com/dgrijalva/jwt-go"

type ClaimsDTO struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}
