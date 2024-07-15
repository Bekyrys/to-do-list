package internal

import (
	"errors"
	"sync"
	"time"
)

var (
	tasks = make(map[string]Task)
	mu    sync.Mutex
)

func GetTasks(status string) []Task {
	mu.Lock()
	defer mu.Unlock()

	var result []Task
	for _, task := range tasks {
		if status == "done" && task.Done {
			result = append(result, task)
		} else if status != "done" && !task.Done && task.ActiveAt.Before(time.Now()) {
			result = append(result, task)
		}
	}
	return result
}

func AddTask(task Task) error {
	mu.Lock()
	defer mu.Unlock()

	for _, t := range tasks {
		if t.Title == task.Title && t.ActiveAt == task.ActiveAt {
			return errors.New("task already exists")
		}
	}

	tasks[task.ID] = task
	return nil
}

func UpdateTask(task Task) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := tasks[task.ID]; !exists {
		return errors.New("task not found")
	}

	tasks[task.ID] = task
	return nil
}

func DeleteTask(id string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(tasks, id)
	return nil
}

func MarkTaskAsDone(id string) error {
	mu.Lock()
	defer mu.Unlock()

	task, exists := tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	task.Done = true
	tasks[id] = task
	return nil
}
