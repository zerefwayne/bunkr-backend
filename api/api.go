package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zerefwayne/college-portal-backend/auth"
	"github.com/zerefwayne/college-portal-backend/config"
	course_http "github.com/zerefwayne/college-portal-backend/course/http"
	resource_http "github.com/zerefwayne/college-portal-backend/resource/http"
	"github.com/zerefwayne/college-portal-backend/user"
	"github.com/zerefwayne/college-portal-backend/utils"
)

// NewRouter defines a mux router and attaches all the routes from various packages.
// Returns a fully configured mux router.
func NewRouter() *mux.Router {

	router := mux.NewRouter()

	// Makes it compulsory to NOT include a slash after a route URL.
	router.StrictSlash(true)
	// Calls the Logger middleware for every route.
	router.Use(utils.Logger)

	// Attach routes from packages

	user.SetUserHandlers(router.PathPrefix("/api/user").Subrouter())
	resource_http.SetResourceHandlers(router.PathPrefix("/api/resource").Subrouter())
	auth.SetAuthHandlers(router.PathPrefix("/api/auth").Subrouter())
	course_http.SetCourseHandlers(router.PathPrefix("/api/course").Subrouter())

	return router
}

// CORSHandler creates a new mux router and allows CORS from all sources
func CORSHandler() http.Handler {

	// router is a fully configured mux Router
	router := NewRouter()

	// attaches AllowAll course policies to router
	corsHandler := cors.AllowAll().Handler(router)

	return corsHandler

}

// getPort generates a port from environment variables.
func getPort() string {

	// Env variable PORT overrides the normal APIPort if available

	if config.C.Env.APIEnv.Port != "" {
		return ":" + config.C.Env.APIEnv.Port
	}

	port := config.C.Env.APIEnv.APIPort
	return ":" + port

}

// Serve starts the REST API on the port
func Serve() {

	handler := CORSHandler()

	log.Printf("api		| listening on port %s\n", getPort())

	log.Fatalln(http.ListenAndServe(getPort(), handler))

}
