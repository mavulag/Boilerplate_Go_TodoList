// cmd/api/main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mavulag/Boilerplate_Go_TodoList/internal/db"
	"github.com/mavulag/Boilerplate_Go_TodoList/internal/handlers"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize PostgreSQL and MongoDB connections
	// postgresDBURI := os.Getenv("POSTGRES_DB_URI")
	// postgresDB, err := db.NewPostgresDB(postgresDBURI)
	// if err != nil {
	// 	log.Fatalf("Error initializing PostgreSQL database: %v", err)
	// }

	mongoDBURI := os.Getenv("MONGO_DB_URI")
	mongoDB, err := db.NewMongoDB(mongoDBURI)
	if err != nil {
		log.Fatalf("Error initializing MongoDB database: %v", err)
	}

	// Initialize router
	router := mux.NewRouter()

	// Initialize handlers with database connections
	// taskHandler := handlers.NewTaskHandler(postgresDB, mongoDB)
	taskHandler := handlers.NewTaskHandler(mongoDB)

	// Define API routes
	router.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", taskHandler.GetTask).Methods("GET")
	router.HandleFunc("/task", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/task/{id}", taskHandler.DeleteTask).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
