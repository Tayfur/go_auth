package routes

import (
	"go_auth/pkg/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	auth := e.Group("/auth")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status},ms=${latency_human}\n",
	}))
	auth.POST("/login", func(c echo.Context) error {
		return handlers.Login(c, db)
	})
	auth.POST("/register", func(c echo.Context) error {
		return handlers.Register(c, db)
	})
	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secureSecretText"),
	}))

	api.GET("/me", func(c echo.Context) error {
		return handlers.Me(c, db)
	})

}
