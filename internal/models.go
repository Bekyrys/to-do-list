package internal

// Task represents a task in the todo list.
type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ActiveAt string `json:"activeAt"`
	Status   string `json:"status"`
}
