package duedatecalc

import (
	"fmt"
	"time"
)

func isValidTurnaround(i int) bool {
	return i > 0
}

func isWeekend(t time.Time) bool {
	switch t.Weekday() {
	case time.Friday:
		h, _, _ := t.Clock()
		if h >= 17 {
			return true
		}
	case time.Saturday:
		return true
	case time.Sunday:
		return true
	}
	return false
}

func isWithinWorkingHours(start string, end string, checkDate time.Time) (bool, error) {
	h, m, _ := checkDate.Clock()
	extractedHours := fmt.Sprintf("%d:%02d", h, m)
	check, err := time.Parse(timeFormat, extractedHours)
	if err != nil {
		return false, fmt.Errorf("parse hours from given date: %w", err)
	}

	workingHoursStart, err := time.Parse(timeFormat, start)
	if err != nil {
		return false, fmt.Errorf("parse working hours start: %w", err)
	}

	workingHoursEnd, err := time.Parse(timeFormat, end)
	if err != nil {
		return false, fmt.Errorf("parse working hours end: %w", err)
	}

	if workingHoursStart.Before(workingHoursEnd) {
		return !check.Before(workingHoursStart) && !check.After(workingHoursEnd), nil
	}
	if workingHoursStart.Equal(workingHoursEnd) {
		return check.Equal(workingHoursStart), nil
	}
	return !workingHoursStart.After(check) || !workingHoursEnd.Before(check), nil
}
