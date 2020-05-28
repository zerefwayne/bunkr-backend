package resource

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/auth"
	"github.com/zerefwayne/college-portal-backend/config"
)

var (
	repo    Repository
	usecase Usecase
)

func SetResourceHandlers(r *mux.Router) {

	repo = NewMongoResourceRepository(config.C.MongoDB)
	usecase = NewResourceUsecase(repo)

	r.Use(auth.SecureRoute)

	r.HandleFunc("/test", defaultHandler)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello from resource!")

}
