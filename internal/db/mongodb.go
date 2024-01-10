// internal/db/mongodb.go
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/mavulag/Boilerplate_Go_TodoList/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	Connection *mongo.Database
}

func NewMongoDB(connection string) (*MongoDB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connection))
	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %w", err)
	}

	fmt.Println("Connected to MongoDB database")

	return &MongoDB{Connection: client.Database("trilabs_todotask")}, nil
}

// GetTasks retrieves all tasks from MongoDB
func (m *MongoDB) GetTasks() ([]models.Task, error) {
	cursor, err := m.Connection.Collection("tasks").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error querying tasks from MongoDB: %w", err)
	}
	defer cursor.Close(context.Background())

	var tasks []models.Task
	for cursor.Next(context.Background()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, fmt.Errorf("error decoding task from MongoDB cursor: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetTask retrieves a task by ID from MongoDB
func (m *MongoDB) GetTask(id string) (*models.Task, error) {
	var task models.Task
	err := m.Connection.Collection("tasks").FindOne(context.Background(), bson.M{"id": id}).Decode(&task)
	if err != nil {
		return nil, fmt.Errorf("error decoding task from MongoDB: %w", err)
	}

	return &task, nil
}

// // CreateTask creates a new task in MongoDB
// func (m *MongoDB) CreateTask(task models.Task) (*models.Task, error) {
// 	_, err := m.Connection.Collection("tasks").InsertOne(context.Background(), task)
// 	if err != nil {
// 		return nil, fmt.Errorf("error inserting task into MongoDB: %w", err)
// 	}

// 	return &task, nil
// }

func (m *MongoDB) CreateTask(title, content string) (*models.Task, error) {
	// Use the NewTask function to generate a new UUID for the task ID
	task := models.NewTask(title, content)

	_, err := m.Connection.Collection("tasks").InsertOne(context.Background(), task)
	if err != nil {
		return nil, fmt.Errorf("error inserting task into MongoDB: %w", err)
	}

	return task, nil
}

// UpdateTask updates a task in MongoDB by ID
func (m *MongoDB) UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedTask}
	_, err := m.Connection.Collection("tasks").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("error updating task in MongoDB: %w", err)
	}

	return &updatedTask, nil
}

// DeleteTask deletes a task by ID from MongoDB
func (m *MongoDB) DeleteTask(id string) error {
	_, err := m.Connection.Collection("tasks").DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("error deleting task from MongoDB: %w", err)
	}

	return nil
}
