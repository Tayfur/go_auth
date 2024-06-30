package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(UserId, Name, Email string) (string, error) {
	claims := Claims{
		UserId: UserId,
		Name:   Name,
		Email:  Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secureSecretText"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
