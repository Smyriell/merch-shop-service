package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("supersecretkey") // Temporary is here

type TockenData struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateTocken(userID int, email string) (string, error) {
	tockenData := TockenData{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)), // Будем пока выдавать токен на 12 часов
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, tockenData)
	return token.SignedString(secretKey)
}
