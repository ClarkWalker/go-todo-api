package main

import (
	"net/http"
)

// OneRoute ...
type OneRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []OneRoute

var routeList = Routes{
	OneRoute{
		"Index",
		"GET",
		"/",
		Index,
	},
	OneRoute{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	OneRoute{
		"TodoShow",
		"GET",
		"/todos/{ID}",
		TodoShow,
	},
}
