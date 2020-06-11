package utils

import (
	"log"
	"net/http"
)

// Logger is used to log the incoming requests URI and method
func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})

}
