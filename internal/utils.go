package internal

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func IsWeekend(date string) bool {
	layout := "2006-01-02"
	d, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	weekday := d.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}
