package routes

import (
	"dev-book-api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Handler:        controllers.CreateUser,
		Description:    "Route to create a new user",
		IsAuthRequired: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Handler:        controllers.GetAllUsers,
		Description:    "Route to list all users",
		IsAuthRequired: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Handler:        controllers.GetUserById,
		Description:    "Route to get user by ID",
		IsAuthRequired: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodPut,
		Handler:        controllers.UpdateUser,
		Description:    "Route to update user",
		IsAuthRequired: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodDelete,
		Handler:        controllers.DeleteUser,
		Description:    "Route to delete user",
		IsAuthRequired: false,
	},
}
