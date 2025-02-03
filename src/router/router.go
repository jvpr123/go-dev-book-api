package router

import "github.com/gorilla/mux"

// Returns a new router with routers setup
func GenerateRouter() *mux.Router {
	return mux.NewRouter()
}
