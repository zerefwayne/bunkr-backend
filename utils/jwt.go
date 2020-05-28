package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
)

func GenerateJWTString(user *models.User) (string, error) {

	signingKey := []byte(config.C.Env.APIEnv.SigningKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 5).Unix(),
		"id":  user.ID,
	})

	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}

func VerifyJWT(tokenString string) (jwt.Claims, error) {

	signingKey := []byte(config.C.Env.APIEnv.SigningKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	return token.Claims, err

}

// SecureRoute secures the route via JWT
func SecureRoute(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if len(tokenString) == 0 {
			http.Error(w, "authentication error: login required", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := VerifyJWT(tokenString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claimsMap := claims.(jwt.MapClaims)

		fmt.Println(claimsMap)

		id := fmt.Sprintf("%v", claimsMap["id"])

		r.Header.Set("id", id)

		next.ServeHTTP(w, r)

	})

}
