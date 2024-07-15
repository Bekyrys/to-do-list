package internal

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

// ValidateTask validates the task fields
func ValidateTask(task Task) error {
	if utf8.RuneCountInString(task.Title) == 0 || utf8.RuneCountInString(task.Title) > 200 {
		return errors.New("title must be between 1 and 200 characters")
	}
	if task.ActiveAt.IsZero() || !task.ActiveAt.After(time.Time{}) {
		return errors.New("invalid activeAt date")
	}
	return nil
}

// GenerateID generates a unique ID for a task
func GenerateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// IsWeekend checks if the given date is a weekend
func IsWeekend(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}
