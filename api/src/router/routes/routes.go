package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Represents all routes
type Route struct {
	Uri          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Config put all routes inside router
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}
