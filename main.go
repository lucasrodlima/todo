package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Todo struct {
	Title       string    `json:"name"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

func createTodo(title string, taskFile string) (Todo, error) {
	newTask := Todo{
		Title:     title,
		Status:    false,
		CreatedAt: time.Now(),
	}

	fmt.Println(newTask.Title)

	data, err := json.Marshal(newTask)
	if err != nil {
		return Todo{}, err
	}

	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		return Todo{}, err
	}

	return newTask, nil
}

func main() {
	path := os.Getenv("TASK_FILE")

	testTask, err := createTodo("Study today", path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(testTask.Title)
}
