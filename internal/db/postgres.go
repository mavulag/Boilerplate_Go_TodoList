// internal/db/postgres.go
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/mavulag/Boilerplate_Go_TodoList/internal/models"
)

type PostgresDB struct {
	Connection *sql.DB
}

func NewPostgresDB(connection string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, fmt.Errorf("error opening PostgreSQL connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging PostgreSQL database: %w", err)
	}

	fmt.Println("Connected to PostgreSQL database")

	return &PostgresDB{Connection: db}, nil
}

func (p *PostgresDB) GetTasks() ([]models.Task, error) {
	// Implementation to retrieve tasks from PostgreSQL
	rows, err := p.Connection.Query("SELECT id, title, content FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("error querying tasks from PostgreSQL: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Content)
		if err != nil {
			return nil, fmt.Errorf("error scanning task row: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (p *PostgresDB) GetTask(id string) (*models.Task, error) {
	// Implementation to retrieve a task by ID from PostgreSQL
	row := p.Connection.QueryRow("SELECT id, title, content FROM tasks WHERE id = $1", id)

	var task models.Task
	err := row.Scan(&task.ID, &task.Title, &task.Content)
	if err != nil {
		return nil, fmt.Errorf("error scanning task row: %w", err)
	}

	return &task, nil
}

func (p *PostgresDB) CreateTask(task models.Task) (*models.Task, error) {
	// Implementation to create a task in PostgreSQL
	_, err := p.Connection.Exec("INSERT INTO tasks (id, title, content) VALUES ($1, $2, $3)", task.ID, task.Title, task.Content)
	if err != nil {
		return nil, fmt.Errorf("error inserting task into PostgreSQL: %w", err)
	}

	return &task, nil
}

func (p *PostgresDB) UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	// Implementation to update a task in PostgreSQL
	_, err := p.Connection.Exec("UPDATE tasks SET title = $1, content = $2 WHERE id = $3", updatedTask.Title, updatedTask.Content, id)
	if err != nil {
		return nil, fmt.Errorf("error updating task in PostgreSQL: %w", err)
	}

	return &updatedTask, nil
}

func (p *PostgresDB) DeleteTask(id string) error {
	// Implementation to delete a task by ID from PostgreSQL
	_, err := p.Connection.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting task from PostgreSQL: %w", err)
	}

	return nil
}



// // internal/db/postgres.go
// package db

// import (
// 	"database/sql"
// 	"fmt"
// 	"os"

// 	"github.com/joho/godotenv"
// 	_ "github.com/lib/pq"
// )

// // For PostgreSQL

// var POSTGRES_DB *sql.DB

// func InitializePostgresDB() (*sql.DB, error) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		return nil, fmt.Errorf("error loading .env file: %w", err)
// 	}

// 	username := os.Getenv("POSTGRES_DB_USERNAME")
// 	databaseName := os.Getenv("POSTGRES_DB_DATABASE")
// 	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", username, databaseName)

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("error opening database connection: %w", err)
// 	}

// 	if err = db.Ping(); err != nil {
// 		return nil, fmt.Errorf("error pinging database: %w", err)
// 	}

// 	fmt.Println("Connected to PostgreSQL database")

// 	POSTGRES_DB = db

// 	return db, nil
// }
