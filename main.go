package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Todo struct {
	Title       string    `json:"name"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

func createTodo(title string, path string) (Todo, error) {
	newTask := Todo{
		Title:     title,
		Status:    false,
		CreatedAt: time.Now(),
	}

	data, err := json.Marshal(newTask)
	if err != nil {
		return Todo{}, err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return Todo{}, err
	}

	return newTask, nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}

	taskFile := os.Getenv("TASK_FILE")

	testTask, err := createTodo("Study today\n", taskFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Test Task: %v", testTask.Title)
}
