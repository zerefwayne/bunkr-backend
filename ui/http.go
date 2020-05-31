package ui

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/course"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/user"
	"github.com/zerefwayne/college-portal-backend/utils"
)

func SetUIHandlers(r *mux.Router) {

	r.Use(utils.SecureRoute)

	r.HandleFunc("/init", initHandler)

}

func respond(w http.ResponseWriter, body interface{}, code int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func initHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("id")

	var initResponse struct {
		Success bool             `json:"success"`
		Error   string           `json:"error,omitempty"`
		User    *models.User     `json:"user,omitempty"`
		Courses []*models.Course `json:"courses,omitempty"`
	}

	user, err := user.UserUsecase.GetByID(context.Background(), userID)
	courses, err := course.CourseUsecase.GetAllCourses(context.Background())

	if err != nil {

		initResponse.Success = false
		initResponse.Error = err.Error()

		respond(w, initResponse, http.StatusInternalServerError)
		return
	}

	initResponse.Success = true
	initResponse.User = user
	initResponse.Courses = courses

	respond(w, initResponse, http.StatusOK)

}
