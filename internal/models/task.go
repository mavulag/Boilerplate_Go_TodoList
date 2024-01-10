// // internal/models/task.go
// package models

// type Task struct {
// 	ID      string `json:"id"`
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// 	// Add more fields as needed
// }
// CREATE TABLE tasks (
//     id UUID DEFAULT gen_random_uuid() NOT NULL,
//     title VARCHAR(255) NOT NULL,
//     content TEXT,
//     PRIMARY KEY (id)
// );

// internal/models/task.go
package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewTask creates a new Task instance with a generated UUID.
func NewTask(title, content string) *Task {
	return &Task{
		ID:        uuid.New().String(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
