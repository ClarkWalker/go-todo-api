package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "route: GET to /\nreturn: all records")
}

// TodoIndex ...
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	println("TodoIndex, GET to /todos")

	todos := Todos{
		Todo{Name: "complete go api"},
		Todo{Name: "complete scoreboard api in go"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow ...
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
