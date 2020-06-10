package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/common"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/resource/usecase"
	"github.com/zerefwayne/college-portal-backend/user"
	"github.com/zerefwayne/college-portal-backend/utils"
)

func SetResourceHandlers(r *mux.Router) {

	common.Resource = usecase.NewResourceUsecase()

	r.Use(utils.SecureRoute)

	r.HandleFunc("/test", defaultHandler)
	r.HandleFunc("/create", createResourceHandler)
	r.HandleFunc("/update", updateResourceHandler)
	r.HandleFunc("/", getResourceHandler)
	r.HandleFunc("/user", getUserResources)
	r.HandleFunc("/all", getAllResources)
	r.HandleFunc("/delete", deleteResourceByIDHandler)
	r.HandleFunc("/pending", pendingResourceHandler)
	r.HandleFunc("/approve", approveResourceByIDHandler)
	r.HandleFunc("/add-vote", addVoteHandler)
	r.HandleFunc("/update-vote", updateVoteHandler)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello from resource!")

}

func addVoteHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("id")
	resourceID := r.URL.Query().Get("resourceID")

	if userID == "" || resourceID == "" {
		log.Println(userID, resourceID)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err := common.Resource.AddVoteResource(context.Background(), resourceID, userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Respond(w, "Successful!", http.StatusOK)

}

func updateVoteHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("id")
	resourceID := r.URL.Query().Get("resourceID")

	if userID == "" || resourceID == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err := common.Resource.UpdateVoteResource(context.Background(), resourceID, userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Respond(w, "Successful!", http.StatusOK)

}

func pendingResourceHandler(w http.ResponseWriter, r *http.Request) {

	resources, err := common.Resource.GetPendingResources(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var payload struct {
		Resources []*models.Resource `json:"resources"`
	}

	payload.Resources = resources

	utils.Respond(w, payload, http.StatusOK)

}

func getResourceHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	resource, err := common.Resource.GetResourceByID(context.Background(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user, err := user.UserUsecase.GetByID(context.Background(), resource.CreatedBy)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var payload struct {
		Resource *models.Resource `json:"resource,omitempty"`
		User     *models.User     `json:"user,omitempty"`
	}

	payload.Resource = resource
	payload.User = user

	utils.Respond(w, payload, http.StatusOK)

}

func createResourceHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("id")

	var body struct {
		Content    string   `json:"content,omitempty"`
		CourseCode string   `json:"courseCode,omitempty"`
		Type       string   `json:"type,omitempty"`
		Title      string   `json:"title,omitempty"`
		Tags       []string `json:"tags,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var newResource models.Resource

	newResource.Content = body.Content
	newResource.CreatedBy = userID
	newResource.Type = body.Type
	newResource.Title = body.Title
	newResource.Tags = body.Tags
	newResource.Upvotes = []string{}

	if err := common.Resource.CreateResource(context.Background(), &newResource); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := common.Course.PushResource(context.Background(), body.CourseCode, newResource.ID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(newResource); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func updateResourceHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("id")

	var updated *models.Resource

	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if updated.CreatedBy != userID {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if err := common.Resource.UpdateResource(context.Background(), updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func getUserResources(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("id")

	resources, err := common.Resource.GetResourcesByUserID(context.Background(), userID)

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

	if len(resources) == 0 {
		response.Resources = []*models.Resource{}
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func getAllResources(w http.ResponseWriter, r *http.Request) {

	resources, err := common.Resource.GetResourcesAll(context.Background())

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
		ID string `json:"resourceID,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err := common.Resource.DeleteResourceByID(context.Background(), body.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "delete success %s", body.ID)

}

func approveResourceByIDHandler(w http.ResponseWriter, r *http.Request) {

	var body struct {
		ID string `json:"id,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err := common.Resource.ApproveResourceByID(context.Background(), body.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload struct {
		Success bool `json:"success,omitempty"`
	}

	payload.Success = true

	utils.Respond(w, payload, http.StatusOK)

}
