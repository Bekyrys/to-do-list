package internal

import (
	"errors"
	"fmt"
	"sync"
)

type InMemoryDB struct {
	tasks map[string]Task
	mu    sync.Mutex
}

var db = InMemoryDB{
	tasks: make(map[string]Task),
}

func AddTask(task Task) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	key := fmt.Sprintf("%s-%s", task.Title, task.ActiveAt)
	if _, exists := db.tasks[key]; exists {
		return errors.New("task already exists")
	}

	db.tasks[task.ID] = task
	return nil
}

func UpdateTask(task Task) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.tasks[task.ID]; !exists {
		return errors.New("task not found")
	}

	db.tasks[task.ID] = task
	return nil
}

func DeleteTask(id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(db.tasks, id)
	return nil
}

func MarkTaskAsDone(id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	task, exists := db.tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	task.Status = "done"
	db.tasks[id] = task
	return nil
}

func GetTasks(status string) []Task {
	db.mu.Lock()
	defer db.mu.Unlock()

	var tasks []Task
	for _, task := range db.tasks {
		if status == "" || task.Status == status {
			tasks = append(tasks, task)
		}
	}

	return tasks
}

func GetTaskByID(id string) (Task, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	task, exists := db.tasks[id]
	if !exists {
		return Task{}, errors.New("task not found")
	}

	return task, nil
}
