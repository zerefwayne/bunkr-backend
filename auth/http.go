package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/user"
	"github.com/zerefwayne/college-portal-backend/utils"
)

type signUpBody struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type loginBody struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// SetAuthHandlers ...
func SetAuthHandlers(r *mux.Router) {

	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/signup", signUpHandler)
}

func respond(w http.ResponseWriter, body interface{}, code int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	var body loginBody

	var loginResponse struct {
		Success bool         `json:"success"`
		Error   string       `json:"error,omitempty"`
		User    *models.User `json:"user,omitempty"`
		Token   string       `json:"token,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {

		loginResponse.Success = false
		loginResponse.Error = err.Error()

		respond(w, loginResponse, http.StatusInternalServerError)
		return

	}

	defer r.Body.Close()

	var userX *models.User
	var err error

	if len(body.Email) > 0 {
		userX, err = user.UserUsecase.GetByEmail(context.Background(), body.Email)
	} else if len(body.Username) > 0 {
		userX, err = user.UserUsecase.GetByUsername(context.Background(), body.Username)
	}

	if userX == nil {

		loginResponse.Success = false
		loginResponse.Error = "user not found"

		respond(w, loginResponse, http.StatusUnauthorized)
		return

	}

	if err != nil {
		loginResponse.Success = false
		loginResponse.Error = "user not found"

		respond(w, loginResponse, http.StatusUnauthorized)
		return
	}

	if err := utils.CompareHashAndPassword(body.Password, userX.Password); err != nil {
		loginResponse.Success = false
		loginResponse.Error = "incorrect password"

		respond(w, loginResponse, http.StatusUnauthorized)
		return
	}

	userX.Password = ""

	jwtToken, err := utils.GenerateJWTString(userX)

	if err != nil {
		loginResponse.Success = false
		loginResponse.Error = "error in jwt creation"

		respond(w, loginResponse, http.StatusInternalServerError)
		return
	}

	loginResponse.Success = true
	loginResponse.Error = ""
	loginResponse.Token = jwtToken

	respond(w, loginResponse, http.StatusOK)

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "/api/auth/logout")
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {

	var body signUpBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	newUser := &models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
	}

	err := user.UserUsecase.CreateUser(context.Background(), newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User creation successful! %+v\n", newUser)
}
