package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/resource"
	"github.com/zerefwayne/college-portal-backend/user"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	router.StrictSlash(true)

	user.SetUserHandlers(router.PathPrefix("/api/user").Subrouter())
	resource.SetResourceHandlers(router.PathPrefix("/api/resource").Subrouter())

	return router
}

func CORSHandler() http.Handler {

	router := NewRouter()

	corsHandler := cors.AllowAll().Handler(router)

	return corsHandler

}

func getPort() string {
	port := config.C.Env.APIEnv.Port
	return ":" + port
}

func Serve() {

	handler := CORSHandler()

	log.Printf("api		| listening on port %s\n", getPort())

	log.Fatalln(http.ListenAndServe(getPort(), handler))

}
