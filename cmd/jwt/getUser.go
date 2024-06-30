package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) (jwt.Claims, error) {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil

}
