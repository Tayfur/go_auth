package models

import (
	_ "github.com/go-playground/validator/v10"
)

type User struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
type LoginUserDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type RegisterUserDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string ` json:"email" validate:"required" gorm:"unique"`
	Password string `json:"password" validate:"required"`
}

type UserRedis struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}
