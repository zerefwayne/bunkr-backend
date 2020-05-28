package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetAuthHandlers(r *mux.Router) {
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/signup", signUpHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "/api/auth/login")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "/api/auth/logout")
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "/api/auth/signup")
}
