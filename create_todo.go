package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func createTodo(args []string) error {
	path := os.Getenv("TASK_FILE")

	if len(args) < 2 {
		return fmt.Errorf("Not enough arguments")
	}

	title := args[1]
	newTask := Todo{
		Title:     title,
		Status:    Pending,
		CreatedAt: time.Now(),
	}

	tasksData, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		tasksData = []byte("[]")
	} else if err != nil {
		return nil
	}

	var tasks []Todo

	err = json.Unmarshal(tasksData, &tasks)
	if err != nil {
		return err
	}

	tasks = append(tasks, newTask)

	newTasksData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, newTasksData, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Task created!: %v\n", newTask.Title)

	return nil
}
