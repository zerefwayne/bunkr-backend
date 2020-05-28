package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/user"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	user user.Usecase
}

var usecase authUsecase

type signUpBody struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func initUsecase() {

	userRepository := user.NewMongoUserRepository(config.C.MongoDB)
	userUsecase := user.NewUserUsecase(userRepository)

	usecase = authUsecase{
		user: userUsecase,
	}

}

func SetAuthHandlers(r *mux.Router) {

	initUsecase()

	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/signup", signUpHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	var body loginBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	user, err := usecase.user.GetByUsername(context.Background(), body.Username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := compareHashAndPassword(body.Password, user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password = ""

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func compareHashAndPassword(password string, hash string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err

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

	err := usecase.user.CreateUser(context.Background(), newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User creation successful! %+v\n", newUser)

}
