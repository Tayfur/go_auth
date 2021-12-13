package handlers

import (
	"errors"
	"go_auth/models"
	"go_auth/pkg/config"
	"go_auth/pkg/jwt"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Register(c echo.Context, db *gorm.DB) error {
	temp := models.RegisterUserDto{}
	validate := validator.New()
	err := c.Bind(&temp)
	if err != nil {
		log.Print(err)
		return c.JSON(400, models.Error{Message: "Invalid field", Status: 400, Error: err.Error()})
	}
	err = validate.Struct(&temp)
	if err != nil {
		return c.JSON(400, models.Error{Message: "Invalid field", Status: 400, Error: err.Error()})
	}
	hashPassword, err := config.GenerateHashPassword(temp.Password)
	if err != nil {
		log.Print(err)
	}
	user := &models.User{
		UserId:   uuid.NewString(),
		Name:     temp.Name,
		Email:    temp.Email,
		Password: hashPassword,
	}
	err = db.Create(user).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.Print(err)
			return c.JSON(500, models.Error{Message: "You are already registered", Status: 500, Error: pgErr.Message})
		}
	}
	token, err := jwt.GenerateToken(user.UserId, user.Name, user.Email)
	if err != nil {
		return c.JSON(500, "Cannot Response Token")
	}

	return c.JSON(200, token)
}
