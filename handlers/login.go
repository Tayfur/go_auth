package handlers

import (
	"go_auth/models"
	"go_auth/pkg/config"
	"go_auth/pkg/jwt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Login(c echo.Context, db *gorm.DB) error {
	userDto := new(models.LoginUserDto)
	err := c.Bind(userDto)
	if err != nil {
		return c.JSON(400, models.Error{
			Message: "Invalid request",
			Error:   err.Error(),
			Status:  400,
		})
	}
	validate := validator.New()
	_ = validate.Struct(userDto)
	user := new(models.User)
	db.Where(models.User{Email: userDto.Email}).First(user)
	checkPassword := config.CheckPasswordHash(userDto.Password, user.Password)
	if !checkPassword {
		return c.JSON(401, models.Error{
			Message: "Invalid password",
			Error:   "Invalid password",
			Status:  401,
		})
	}
	token, err := jwt.GenerateToken(user.UserId, user.Name, user.Email)
	if err != nil {
		return c.JSON(500, models.Error{
			Message: "Cannot Return Token",
			Error:   err.Error(),
			Status:  500,
		})
	}
	return c.JSON(200, token)
}
