package handlers

import (
	"go_auth/pkg/jwt"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Me(c echo.Context, db *gorm.DB, redis *redis.Client) error {
	user, _ := jwt.GetUser(c)
	return c.JSON(200, user)
}
