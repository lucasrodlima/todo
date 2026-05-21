package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func GetTasks() ([]*Todo, error) {

	path := os.Getenv("TASK_FILE")

	var tasks []*Todo

	tasksData, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		tasksData = []byte("[]")
	} else if err != nil {
		return nil, err
	}

	err = json.Unmarshal(tasksData, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTask(newTask *Todo) error {
	path := os.Getenv("TASK_FILE")

	tasks, err := GetTasks()
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

func ChangeStatus(completeId int64) error {
	todos, err := GetTasks()
	if err != nil {
		return err
	}

	for _, t := range todos {
		if t.Id == completeId {
			if t.Status == Done {
				return fmt.Errorf("Task is already completed!")
			}

			t.Status = Done
			fmt.Printf("Task %v - %v Completed!\n", t.Id, t.Title)

			now := time.Now()
			t.CompletedAt = now
		}
	}

	newTasksData, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}

	path := os.Getenv("TASK_FILE")
	err = os.WriteFile(path, newTasksData, 0644)
	if err != nil {
		return err
	}

	return nil
}
