package ui

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/course"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/user"
	"github.com/zerefwayne/college-portal-backend/utils"
)

type uiUsecase struct {
	user   user.Usecase
	course course.Usecase
}

var usecase uiUsecase

func initUsecase() {

	userRepository := user.NewMongoUserRepository(config.C.MongoDB)
	userUsecase := user.NewUserUsecase(userRepository)

	courseUsecase := course.NewCourseUsecase()

	usecase = uiUsecase{
		user:   userUsecase,
		course: courseUsecase,
	}

}

func SetUIHandlers(r *mux.Router) {

	initUsecase()

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

	user, err := usecase.user.GetByID(context.Background(), userID)
	courses, err := usecase.course.GetAllCourses(context.Background())

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
