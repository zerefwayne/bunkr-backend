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

// GenerateJWTString takes in user as input and generates a JWT Token to be supplied with login request.
func GenerateJWTString(user *models.User) (string, error) {

	signingKey := []byte(config.C.Env.APIEnv.SigningKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"id":  user.ID,
	})

	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}

// VerifyJWT parses the JWT Token and returns the claims stored in it
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

		// Reads the Authorization Header from the request
		tokenString := r.Header.Get("Authorization")

		// returns 401 if token string is empty
		if len(tokenString) == 0 {
			http.Error(w, "authentication error: login required", http.StatusUnauthorized)
			return
		}

		// Replaces 'Bearer ' with empty string
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// the token is decrypted into claims
		claims, err := VerifyJWT(tokenString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// claims is converted to a map object
		claimsMap := claims.(jwt.MapClaims)

		id := fmt.Sprintf("%v", claimsMap["id"])

		// A new field 'id' is set in the Header for further processing
		r.Header.Set("id", id)

		// Request is forwarded from middleware
		next.ServeHTTP(w, r)

	})

}
