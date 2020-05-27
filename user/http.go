package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/college-portal-backend/config"
)

var (
	repo    Repository
	usecase Usecase
)

func SetUserHandlers(r *mux.Router) {

	repo = NewMongoUserRepository(config.C.MongoDB)
	usecase = NewUserUsecase(repo)

	r.HandleFunc("/test", defaultHandler)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello from user!")

}
