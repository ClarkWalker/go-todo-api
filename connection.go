package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connection ...
func Connection() {
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
}
