package echo

import (
	"net/http"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/auth"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const jwtContextKey = "user"

func jwtMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ContextKey: jwtContextKey,
		SigningKey: auth.JWTSigningKey(),
	})
}

func jwtClaims(c echo.Context) (jwt.MapClaims, error) {
	token, ok := c.Get(jwtContextKey).(*jwt.Token)
	if !ok || token == nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid jwt")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid jwt claims")
	}

	return claims, nil
}
