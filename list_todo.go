package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func listTodos(args []string) error {
	path := os.Getenv("TASK_FILE")

	if len(args) != 1 {
		return fmt.Errorf("Wrong argument number!")
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

	if len(tasks) == 0 {
		fmt.Println("No tasks have been added.")
		return nil
	}

	fmt.Println("TASKS:")

	for _, task := range tasks {
		switch task.Status {
		case Pending:
			fmt.Printf("%v - %v | PENDING | Created At: %v\n", task.Id, task.Title, task.CreatedAt)

		case Done:
			fmt.Printf("%v - %v | DONE | Created At: %v | Completed At: %v\n", task.Id, task.Title, task.CreatedAt, task.CompletedAt)
		}
	}

	return nil
}
