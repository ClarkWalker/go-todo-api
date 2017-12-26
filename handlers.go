package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "route: GET to /\nreturn: all records")
}

// TodoIndex ...
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	println("TodoIndex, GET to /todos")
	dbName := os.Getenv("DATABASE_URL")
	if dbName == "" {
		dbName = "host=localhost dbname=golang-todo-api sslmode=disable"
	}
	fmt.Println("database connection: ", dbName)

	db, err := gorm.Open("postgres", dbName)
	if err != nil {
		panic("Oh nopers! U'r database is not finding :o")
	}
	defer db.Close()

	var todos []Todos
	var getAllTodos = db.Find(&todos)
	// todos := Todos{
	// 	Todo{Name: "complete go api"},
	// 	Todo{Name: "complete scoreboard api in go"},
	// }

	// w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(getAllTodos); err != nil {
		panic(err)
	}
}

// TodoShow ...
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
