package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

func completeTodo(args []string) error {
	path := os.Getenv("TASK_FILE")

	if len(args) < 2 {
		return fmt.Errorf("Not enough arguments")
	}

	complete_id, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return err
	}

	var tasks []*Todo

	tasksData, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		tasksData = []byte("[]")
	} else if err != nil {
		return nil
	}

	err = json.Unmarshal(tasksData, &tasks)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.Id == complete_id {
			task.Status = Done
			fmt.Printf("Task %v - %v Completed!\n", task.Id, task.Title)

			now := time.Now()
			task.CompletedAt = now
		}
	}

	newTasksData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, newTasksData, 0644)
	if err != nil {
		return err
	}

	return nil
}
