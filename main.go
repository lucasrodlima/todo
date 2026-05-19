package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type TaskStatus int

const (
	Pending TaskStatus = iota
	Done
)

type Todo struct {
	Title       string     `json:"name"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt time.Time  `json:"completed_at"`
}

// Create enum with pending and completed for status

func createTodo(title string, path string) (*Todo, error) {
	newTask := Todo{
		Title:     title,
		Status:    Pending,
		CreatedAt: time.Now(),
	}

	tasksData, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		tasksData = []byte("[]")
	} else if err != nil {
		return &Todo{}, nil
	}

	var tasks []Todo

	err = json.Unmarshal(tasksData, &tasks)
	if err != nil {
		return &Todo{}, err
	}

	tasks = append(tasks, newTask)

	newTasksData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return &Todo{}, err
	}

	err = os.WriteFile(path, newTasksData, 0644)
	if err != nil {
		return &Todo{}, err
	}

	return &newTask, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}

	path := os.Getenv("TASK_FILE")

	args := os.Args
	if len(args) != 3 {
		log.Fatal("Incorrect number of arguments")
	}

	command := args[1]
	title := args[2]

	switch command {
	case "add":
		newTask, err := createTodo(title, path)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Task created!: %v\n", newTask.Title)
	}
}
