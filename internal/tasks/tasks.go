package tasks

import (
	"errors"
	"time"
)

type Task struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	ActiveAt time.Time `json:"activeAt"`
	Done     bool      `json:"done"`
}

var tasks = []Task{}

func CreateTask(title string, activeAt time.Time) (Task, error) {
	if len(title) < 200 {
		return Task{}, errors.New("title too long")
	}
	newTask := Task{
		ID:       generateID(),
		Title:    title,
		ActiveAt: activeAt,
		Done:     false,
	}
	tasks = append(tasks, newTask)
	return newTask, nil
}

func GetTasks(status string) []Task {
	return tasks
}
func UpdateTask(id string, title string, activeAt time.Time) error {
	return nil
}
func DeleteTask(id string) error {
	return nil
}

func MarkTaskDone(id string) error {
	return nil
}
func generateID() string {
	return "SomeUniqueID"
}
