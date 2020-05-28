package resource

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

var (
	repo    Repository
	usecase Usecase
)

func SetResourceHandlers(r *mux.Router) {

	repo = NewMongoResourceRepository(config.C.MongoDB)
	usecase = NewResourceUsecase(repo)

	r.Use(utils.SecureRoute)

	r.HandleFunc("/test", defaultHandler)
	r.HandleFunc("/create", createResourceHandler)
	r.HandleFunc("/user", getUserResources)
	r.HandleFunc("/all", getAllResources)
	r.HandleFunc("/delete", deleteResourceByIDHandler)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello from resource!")

}

func createResourceHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("id")

	var body struct {
		Content string `json:"content,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	newResource := new(models.Resource)

	newResource.Content = body.Content
	newResource.CreatedBy = userID

	if err := usecase.CreateResource(context.Background(), newResource); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(newResource); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func getUserResources(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("id")

	resources, err := usecase.GetResourcesByUserID(context.Background(), userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var response struct {
		Length    int                `json:"length"`
		Resources []*models.Resource `json:"resources"`
	}

	response.Length = len(resources)
	response.Resources = resources

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func getAllResources(w http.ResponseWriter, r *http.Request) {

	resources, err := usecase.GetResourcesAll(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var response struct {
		Length    int                `json:"length"`
		Resources []*models.Resource `json:"resources"`
	}

	response.Length = len(resources)
	response.Resources = resources

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func deleteResourceByIDHandler(w http.ResponseWriter, r *http.Request) {

	var body struct {
		ID string `json:"id,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err := usecase.DeleteResourceByID(context.Background(), body.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "delete success %s", body.ID)

}
