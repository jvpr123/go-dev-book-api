package router

import (
	"dev-book-api/src/router/routes"

	"github.com/gorilla/mux"
)

// Returns a new router with routers setup
func GenerateRouter() *mux.Router {
	return routes.ConfigRoutes(mux.NewRouter())
}
