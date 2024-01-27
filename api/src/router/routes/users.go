package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		Uri:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		AuthRequired: false,
	},
	{
		Uri:          "/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		AuthRequired: false,
	},
	{
		Uri:          "/users/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		Uri:          "/users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: false,
	},
}
