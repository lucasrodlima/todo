package task

import "time"

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
	Id          int64      `json:"id"`
}
