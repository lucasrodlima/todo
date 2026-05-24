package cli

import (
	"fmt"

	"github.com/lucasrodlima/todo/internal/task"
)

func List(args []string) error {
	var filter string
	if len(args) > 2 {
		return fmt.Errorf("Wrong argument number!")
	} else if len(args) == 2 {
		if args[1] != "--done" && args[1] != "--pending" {
			return fmt.Errorf("Flag does not exist! Use --done or --pending")
		}
		filter = args[1]
	}

	todos, err := task.GetTasks()
	if err != nil {
		return err
	}

	if len(todos) == 0 {
		fmt.Println("No tasks have been added.")
		return nil
	}

	for _, t := range todos {
		switch t.Status {
		case task.Pending:
			if filter != "--done" {
				fmt.Printf("\u274C %v - %v | Created At: %v\n", t.Id, t.Title, t.CreatedAt.Format("02 Jan 2006 15:04"))
			}

		case task.Done:
			if filter != "--pending" {
				fmt.Printf("\u2705 %v - %v | Created At: %v | Completed At: %v\n", t.Id, t.Title, t.CreatedAt.Format("02 Jan 2006 15:04"), t.CompletedAt.Format("02 Jan 2006 15:04"))
			}
		}
	}

	return nil

}
