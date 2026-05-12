package main

import (
	"fmt"
	"time"
)

type Todo struct {
	Title       string    `json:"name"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

func newTodo(title string) {
	newTask := Todo{
		Title:     title,
		Status:    false,
		CreatedAt: time.Now(),
	}

	fmt.Println(newTask.Title)
}

func main() {
	car := "Honda"
	fmt.Printf("Hello %v\n", car)
}
