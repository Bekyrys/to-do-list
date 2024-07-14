package internal

import (
	"errors"
	"regexp"
)

func ValidateTask(task Task) error {
	if len(task.Title) == 0 || len(task.Title) > 200 {
		return errors.New("title must be between 1 and 200 characters")
	}

	datePattern := `^\d{4}-\d{2}-\d{2}$`
	matched, _ := regexp.MatchString(datePattern, task.ActiveAt)
	if !matched {
		return errors.New("activeAt must be a valid date in format YYYY-MM-DD")
	}

	return nil
}

