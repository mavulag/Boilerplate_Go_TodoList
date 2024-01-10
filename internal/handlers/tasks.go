// internal/handlers/tasks.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mavulag/Boilerplate_Go_TodoList/internal/db"
	"github.com/mavulag/Boilerplate_Go_TodoList/internal/models"
)

type TaskHandler struct {
	// postgresDB *db.PostgresDB
	mongoDB    *db.MongoDB
}

// func NewTaskHandler(postgresDB *db.PostgresDB, mongoDB *db.MongoDB) *TaskHandler {
// 	return &TaskHandler{
// 		postgresDB: postgresDB,
// 		mongoDB:    mongoDB,
// 	}
// }
func NewTaskHandler(mongoDB *db.MongoDB) *TaskHandler {
	return &TaskHandler{
		mongoDB:    mongoDB,
	}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// Retrieve tasks from PostgreSQL
	// postgresTasks, err := h.postgresDB.GetTasks()
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error retrieving tasks from PostgreSQL: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	// Retrieve tasks from MongoDB
	mongoTasks, err := h.mongoDB.GetTasks()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving tasks from MongoDB: %v", err), http.StatusInternalServerError)
		return
	}

	// Combine and return tasks
	// allTasks := append(postgresTasks, mongoTasks...)
	allTasks := mongoTasks
	respondJSON(w, allTasks)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from the request parameters
	vars := mux.Vars(r)
	taskID := vars["id"]

	// Retrieve task from PostgreSQL
	// postgresTask, err := h.postgresDB.GetTask(taskID)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error retrieving task from PostgreSQL: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	// Retrieve task from MongoDB
	mongoTask, err := h.mongoDB.GetTask(taskID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving task from MongoDB: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the first non-nil task
	// if postgresTask != nil {
	// 	respondJSON(w, postgresTask)
	// } else 
	if mongoTask != nil {
		respondJSON(w, mongoTask)
	} else {
		http.NotFound(w, r)
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON payload to a Task struct
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	// Create task in PostgreSQL
	// postgresTask, err := h.postgresDB.CreateTask(newTask)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error creating task in PostgreSQL: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	// Create task in MongoDB
	mongoTask, err := h.mongoDB.CreateTask(newTask.Title, newTask.Content)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating task in MongoDB: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the first created task
	// if postgresTask != nil {
	// 	respondJSON(w, postgresTask)
	// } else 
	if mongoTask != nil {
		respondJSON(w, mongoTask)
	} else {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from the request parameters
	vars := mux.Vars(r)
	taskID := vars["id"]

	// Decode the incoming JSON payload to a Task struct
	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	// Update task in PostgreSQL
	// postgresTask, err := h.postgresDB.UpdateTask(taskID, updatedTask)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error updating task in PostgreSQL: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	// Update task in MongoDB
	mongoTask, err := h.mongoDB.UpdateTask(taskID, updatedTask)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating task in MongoDB: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the first updated task
	// if postgresTask != nil {
	// 	respondJSON(w, postgresTask)
	// } else 
	if mongoTask != nil {
		respondJSON(w, mongoTask)
	} else {
		http.NotFound(w, r)
	}
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from the request parameters
	vars := mux.Vars(r)
	taskID := vars["id"]

	// Delete task in PostgreSQL
	// err := h.postgresDB.DeleteTask(taskID)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error deleting task in PostgreSQL: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	// Delete task in MongoDB
	err := h.mongoDB.DeleteTask(taskID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting task in MongoDB: %v", err), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusNoContent)
}

// respondJSON is a helper function to respond with JSON data
func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
