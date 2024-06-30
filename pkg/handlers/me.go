package handlers

import (
	"go_auth/cmd/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Me(c echo.Context, db *gorm.DB) error {
	user, _ := jwt.GetUser(c)
	return c.JSON(200, user)
}
