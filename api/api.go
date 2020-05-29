package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zerefwayne/college-portal-backend/auth"
	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/resource"
	"github.com/zerefwayne/college-portal-backend/ui"
	"github.com/zerefwayne/college-portal-backend/user"
	"github.com/zerefwayne/college-portal-backend/utils"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	router.StrictSlash(true)
	router.Use(utils.Logger)

	user.SetUserHandlers(router.PathPrefix("/api/user").Subrouter())
	resource.SetResourceHandlers(router.PathPrefix("/api/resource").Subrouter())
	auth.SetAuthHandlers(router.PathPrefix("/api/auth").Subrouter())
	ui.SetUIHandlers(router.PathPrefix("/api/ui").Subrouter())

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
