package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Base application route structure
type Route struct {
	URI            string
	Method         string
	Description    string
	Handler        func(http.ResponseWriter, *http.Request)
	IsAuthRequired bool
}

// Load all application routes inside router
func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return r
}
