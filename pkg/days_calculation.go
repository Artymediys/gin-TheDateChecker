package pkg

import (
	"fmt"
	"time"
)

func CalcDays(userYear int) string {
	var currentDate = time.Now()
	var userDate = time.Date(userYear, time.January, 1, 0, 0, 0, 0, time.UTC)

	switch {
	case userDate.Year() < currentDate.Year():
		return fmt.Sprintf("Days gone: %d", int64(time.Since(userDate).Hours())/24)
	case userDate.Year() > currentDate.Year():
		return fmt.Sprintf("Days left: %d", int64(time.Until(userDate).Hours())/24)
	default:
		return fmt.Sprintf("Congrats! You are in %d already!", userDate.Year())
	}
}
