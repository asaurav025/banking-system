package authentication

import (
	"banking-system/pkg/jwt"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type jwtExtractor func(echo.Context) (string, error)

func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", echo.NewHTTPError(http.StatusBadRequest, "missing or malinformed jwt")
	}
}

func AuthenticationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := "Authorization"
			extractor := jwtFromHeader(authHeader, "Bearer")
			auth, err := extractor(c)
			if err != nil {
				return err
			}

			userId, err := jwt.Validate(auth)

			if err != nil {
				log.Error("Failed to validate token. Error: ", err)
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			c.Set("user.id", userId)
			return next(c)
		}
	}
}
