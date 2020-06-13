package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/user"
	"github.com/zerefwayne/college-portal-backend/utils"
)

// SetAuthHandlers sets authentication routes after /api/auth
func SetAuthHandlers(r *mux.Router) {

	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/signup", signUpHandler)

	// /validate is a secured route hence we create a seperate subrouter to attach middleware.
	validate := r.PathPrefix("/validate").Subrouter()

	// attaches secureRoute middleware to /api/auth/validate
	validate.Use(utils.SecureRoute)

	validate.HandleFunc("/", validateHandler)

}

// validateHandler GET /api/auth/validate
// Returns user object if found
func validateHandler(w http.ResponseWriter, r *http.Request) {

	// extracts userID from Header
	id := r.Header.Get("id")

	// Gets user from user usecase
	user, err := user.UserUsecase.GetByID(context.Background(), id)

	if err != nil {
		log.Println("error in validate", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// User exists in we will send updated user

	var payload struct {
		User *models.User `json:"user"`
	}

	payload.User = user

	utils.Respond(w, payload, http.StatusOK)

}

// loginHandler POST /api/auth/login
func loginHandler(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Email    string `json:"email,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}

	var loginResponse struct {
		Success bool         `json:"success"`
		Error   string       `json:"error,omitempty"`
		User    *models.User `json:"user,omitempty"`
		Token   string       `json:"token,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {

		loginResponse.Success = false
		loginResponse.Error = err.Error()

		utils.Respond(w, loginResponse, http.StatusInternalServerError)
		return

	}

	defer r.Body.Close()

	var userX *models.User
	var err error

	// The login response can contain either mail aur username

	if len(body.Email) > 0 {
		userX, err = user.UserUsecase.GetByEmail(context.Background(), body.Email)
	} else if len(body.Username) > 0 {
		userX, err = user.UserUsecase.GetByUsername(context.Background(), body.Username)
	}

	if userX == nil {

		loginResponse.Success = false
		loginResponse.Error = "user not found"

		utils.Respond(w, loginResponse, http.StatusUnauthorized)
		return

	}

	if err != nil {
		loginResponse.Success = false
		loginResponse.Error = "user not found"

		utils.Respond(w, loginResponse, http.StatusUnauthorized)
		return
	}

	// checks if password is valid
	if err := utils.CompareHashAndPassword(body.Password, userX.Password); err != nil {
		loginResponse.Success = false
		loginResponse.Error = "incorrect password"

		utils.Respond(w, loginResponse, http.StatusUnauthorized)
		return
	}

	// unsets user password so it doesn't go in JSON response
	userX.Password = ""

	// generates a JWT Token for the user
	jwtToken, err := utils.GenerateJWTString(userX)

	if err != nil {
		loginResponse.Success = false
		loginResponse.Error = "error in jwt creation"

		utils.Respond(w, loginResponse, http.StatusInternalServerError)
		return
	}

	loginResponse.Success = true
	loginResponse.Error = ""
	loginResponse.User = userX
	loginResponse.Token = jwtToken

	utils.Respond(w, loginResponse, http.StatusOK)

}

// logoutHandler POSt /api/auth/logout
// NOT USED
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "/api/auth/logout")
}

// signUpHandler POST /api/auth/signup
func signUpHandler(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// Defines a new user object
	newUser := &models.User{
		Username:         body.Username,
		Email:            body.Email,
		Password:         body.Password,
		Name:             body.Name,
		IsVerified:       false,
		VerificationCode: xid.New().String(),
		Bookmarks:        []string{},
	}

	// Creates the user
	err := user.UserUsecase.CreateUser(context.Background(), newUser)

	// Error in user creation
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Sending verification code to %s with code %s for user %s\n", newUser.Email, newUser.VerificationCode, newUser.ID)

	response, err := utils.SendVerificationEmail(newUser.VerificationCode, newUser.Name, newUser.Email)

	if err != nil {
		log.Printf("Failed to send verification email %v\n", err)
	} else {
		log.Printf("Successfully send verification email %+v\n", response)
	}

	// User successfully created

	fmt.Fprintf(w, "User creation successful! %+v\n", newUser)
}
