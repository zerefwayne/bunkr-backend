package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/course"
	"github.com/zerefwayne/college-portal-backend/course/usecase"
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

var http_usecase struct {
	course course.Usecase
}

func SetCourseHandlers(r *mux.Router) {

	http_usecase.course = usecase.NewCourseUsecase()

	r.Use(utils.SecureRoute)

	r.HandleFunc("/all", getAllCoursesHandler)
	r.HandleFunc("/", getCourseHandler)
	r.HandleFunc("/new", createCourseHandler)

}

func createCourseHandler(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Slug string `json:"slug,omitempty"`
		Name string `json:"name,omitempty"`
		Code string `json:"code,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	var course models.Course

	course.Code = body.Code
	course.Name = body.Name
	course.Slug = body.Slug

	log.Printf("course created %+v\n", course)

	if err := http_usecase.course.CreateCourse(context.Background(), &course); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload struct {
		Course models.Course `json:"course,omitempty"`
	}

	payload.Course = course

	utils.Respond(w, payload, http.StatusOK)

}

func getAllCoursesHandler(w http.ResponseWriter, r *http.Request) {

	courses, err := http_usecase.course.GetAllCourses(context.Background())

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

	code := r.URL.Query().Get("code")

	course, err := http_usecase.course.GetCourseByCode(context.Background(), code)

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
