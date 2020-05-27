package resource

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetResourceHandlers(r *mux.Router) {

	r.HandleFunc("/test", defaultHandler)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello from resource!")

}
