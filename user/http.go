package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/config"
)

var (
	repo    Repository
	usecase Usecase
)

// SetUserHandlers ...
func SetUserHandlers(r *mux.Router) {

	repo = NewMongoUserRepository(config.C.MongoDB)
	usecase = NewUserUsecase(repo)

	r.HandleFunc("/test", defaultHandler)
	r.HandleFunc("/", getUserByUsernameHandler)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello from user!")

}

func getUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {

	user, err := usecase.GetByUsername(context.Background(), "zerefwayne")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%+v\n", user)

}
