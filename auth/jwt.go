package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zerefwayne/college-portal-backend/models"
)

func generateJWTString(user *models.User) (string, error) {

	signingKey := []byte("secret123")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 5).Unix(),
		"id":  user.ID,
	})

	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}
