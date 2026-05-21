package cli

import (
	"fmt"

	"github.com/lucasrodlima/todo/internal/task"
)

func List(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Wrong argument number!")
	}

	todos, err := task.GetTasks()
	if err != nil {
		return err
	}

	if len(todos) == 0 {
		fmt.Println("No tasks have been added.")
		return nil
	}

	fmt.Println("TASKS:")

	for _, t := range todos {
		switch t.Status {
		case task.Pending:
			fmt.Printf("%v - %v | PENDING | Created At: %v\n", t.Id, t.Title, t.CreatedAt)

		case task.Done:
			fmt.Printf("%v - %v | DONE | Created At: %v | Completed At: %v\n", t.Id, t.Title, t.CreatedAt, t.CompletedAt)
		}
	}

	return nil

}
