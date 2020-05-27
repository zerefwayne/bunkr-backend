package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUserHandlers(r *mux.Router) {

	r.HandleFunc("/test", defaultHandler)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello from user!")

}
