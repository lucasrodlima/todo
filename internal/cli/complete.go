package cli

import (
	"fmt"
	"github.com/lucasrodlima/todo/internal/task"
	"strconv"
)

func Complete(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("Not enough arguments")
	}

	complete_id, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return err
	}

	if err := task.ChangeStatus(complete_id); err != nil {
		return err
	}

	return nil
}
