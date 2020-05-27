package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	return router

}

func CORSHandler() http.Handler {

	router := NewRouter()

	corsHandler := cors.AllowAll().Handler(router)

	return corsHandler

}

func Serve() {

	handler := CORSHandler()

	log.Println("api		| listening on port 5000")

	log.Fatalln(http.ListenAndServe(":5000", handler))

}
