// handlers/handler.go
package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// // Task represents a task entity
// type Task struct {
// 	ID    int    `json:"id"`
// 	Title string `json:"title"`
// }

// // GetTasksHandler retrieves all tasks from both PostgreSQL and MongoDB
// func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
// 	// Fetch tasks from PostgreSQL
// 	postgreSQLTasks, err := pkg.GetPostgreSQLTasks()
// 	if err != nil {
// 		log.Println("Error retrieving tasks from PostgreSQL:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Fetch tasks from MongoDB
// 	mongoDBTasks, err := pkg.GetMongoDBTasks()
// 	if err != nil {
// 		log.Println("Error retrieving tasks from MongoDB:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Combine tasks from both databases
// 	allTasks := append(postgreSQLTasks, mongoDBTasks...)

// 	// Return tasks as JSON
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(allTasks)
// }

// // CreateTaskHandler creates a new task in both PostgreSQL and MongoDB
// func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	var task Task
// 	err := json.NewDecoder(r.Body).Decode(&task)
// 	if err != nil {
// 		log.Println("Error decoding request body:", err)
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	// Create task in PostgreSQL
// 	err = pkg.CreatePostgreSQLTask(task)
// 	if err != nil {
// 		log.Println("Error creating task in PostgreSQL:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Create task in MongoDB
// 	err = pkg.CreateMongoDBTask(task)
// 	if err != nil {
// 		log.Println("Error creating task in MongoDB:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintf(w, "Task created successfully")
// }

// // UpdateTaskHandler updates a task in both PostgreSQL and MongoDB
// func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		log.Println("Error converting task ID to integer:", err)
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	var updatedTask Task
// 	err = json.NewDecoder(r.Body).Decode(&updatedTask)
// 	if err != nil {
// 		log.Println("Error decoding request body:", err)
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	// Update task in PostgreSQL
// 	err = pkg.UpdatePostgreSQLTask(taskID, updatedTask)
// 	if err != nil {
// 		log.Println("Error updating task in PostgreSQL:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Update task in MongoDB
// 	err = pkg.UpdateMongoDBTask(taskID, updatedTask)
// 	if err != nil {
// 		log.Println("Error updating task in MongoDB:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Task updated successfully")
// }

// // DeleteTaskHandler deletes a task from both PostgreSQL and MongoDB
// func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		log.Println("Error converting task ID to integer:", err)
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	// Delete task from PostgreSQL
// 	err = pkg.DeletePostgreSQLTask(taskID)
// 	if err != nil {
// 		log.Println("Error deleting task from PostgreSQL:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Delete task from MongoDB
// 	err = pkg.DeleteMongoDBTask(taskID)
// 	if err != nil {
// 		log.Println("Error deleting task from MongoDB:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Task deleted successfully")
// }
