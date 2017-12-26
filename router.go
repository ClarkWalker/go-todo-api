package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter ...
func NewRouter() *mux.Router {
	var router = mux.NewRouter().StrictSlash(true)
	for _, route := range routeList {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
