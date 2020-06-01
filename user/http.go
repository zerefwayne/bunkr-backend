package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/utils"
)

// SetUserHandlers ...
func SetUserHandlers(r *mux.Router) {

	UserUsecase.userRepo = newMongoUserRepository(config.C.MongoDB)

	r.HandleFunc("/test", defaultHandler)
	r.HandleFunc("/", getUserHandler)

	addCourse := r.PathPrefix("/course").Subrouter()

	addCourse.Use(utils.SecureRoute)

	addCourse.HandleFunc("/add", addCourseHandler)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello from user!")

}

func getUserHandler(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")

	var user *models.User
	var err error

	if len(username) > 0 {
		user, err = UserUsecase.GetByUsername(context.Background(), username)
	} else if len(email) > 0 {
		user, err = UserUsecase.GetByEmail(context.Background(), email)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password = ""

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func addCourseHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("id")

	courseID := r.URL.Query().Get("courseCode")

	if err := UserUsecase.AddCourse(context.Background(), userID, courseID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Respond(w, "success", http.StatusOK)
}
