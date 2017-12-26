package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Todo ...
type Todo struct {
	gorm.Model
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos ...
type Todos []Todo

// Migrations ...
func Migrations() {
	dbName := os.Getenv("DATABASE_URL")
	if dbName == "" {
		dbName = "host=localhost dbname=golang-todo-api sslmode=disable"
	}
	fmt.Println("--migrations.go dbName =", dbName)

	db, err := gorm.Open("postgres", dbName)
	if err != nil {
		panic("Oh noos you got a database nopers")
	}
	defer db.Close()

	db.DropTable(&Todo{})
	db.AutoMigrate(&Todo{})
	db.Create(&Todo{Name: "make go api"})
	db.Create(&Todo{Name: "make it a CRUD api"})
	db.Create(&Todo{Name: "take out the trash"})
} // end Migrations
