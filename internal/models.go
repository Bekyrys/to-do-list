package internal

import "time"

type Task struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	ActiveAt time.Time `json:"activeAt"`
	Done     bool      `json:"done"`
}
