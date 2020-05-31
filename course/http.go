package course

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/utils"
)

func respond(w http.ResponseWriter, body interface{}, code int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func SetCourseHandlers(r *mux.Router) {

	r.Use(utils.SecureRoute)

	r.HandleFunc("/all", getAllCoursesHandler)
	r.HandleFunc("/", getCourseHandler)

}

func getAllCoursesHandler(w http.ResponseWriter, r *http.Request) {

	courses, err := CourseUsecase.GetAllCourses(context.Background())

	if err != nil {
		log.Println("Error while get all courses", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload struct {
		Courses []*models.Course `json:"courses"`
	}

	payload.Courses = courses

	respond(w, payload, http.StatusOK)

}

func getCourseHandler(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("id")

	course, err := CourseUsecase.GetCourseByCode(context.Background(), code)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload struct {
		Course *models.Course `json:"course"`
	}

	payload.Course = course

	respond(w, payload, http.StatusOK)

}
