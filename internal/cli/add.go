package cli

import (
	"fmt"
	"github.com/lucasrodlima/todo/internal/task"
	"time"
)

func Add(args []string) error {
	var taskNumber int64

	if len(args) < 2 {
		return fmt.Errorf("Not enough arguments")
	}

	tasks, err := task.GetTasks()
	if err != nil {
		return err
	}

	if len(tasks) >= 1 {
		taskNumber = int64(len(tasks)) + 1
	} else {
		taskNumber = 1
	}

	title := args[1]

	newTask := task.Todo{
		Title:     title,
		Status:    task.Pending,
		CreatedAt: time.Now(),
		Id:        taskNumber,
	}

	if err := task.SaveTask(&newTask); err != nil {
		return err
	}

	return nil
}
