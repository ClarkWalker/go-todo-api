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
	fmt.Println("migrations db", dbName)

	db, err := gorm.Open("postgres", dbName)
	if err != nil {
		panic("Oh noos you got a database nopers")
	}
	defer db.Close()

	// clear db
	db.DropTable(&Todos{})
	// Migrate the schema
	db.AutoMigrate(&Todos{})
	// Seed the table
	db.Create(&Todos{
		Name: "complete go api"
		// Todo{Name: "complete scoreboard api in go"},
	})

} // end Migrations
