package duedatecalc

import (
	"fmt"
	"time"
)

const dateFormat = "2006-01-02 15:04"
const timeFormat = "15:04"

func DueDateCalculator(issueStart string, turnAround int) (string, error) {
	t, err := time.Parse(dateFormat, issueStart)
	if err != nil {
		return "", fmt.Errorf("parse issue start: %w", err)
	}

	if isWeekend(t) {
		return "", &InputWeekendErr{}
	}

	b, err := isWithinWorkingHours("09:00", "17:00", t)

	if err != nil {
		return "", fmt.Errorf("check input is within working hours")
	}
	if !b {
		return "", &NotWithinWorkingHoursErr{}
	}

	newTime := t.Add(time.Hour * 8)
	return newTime.Format(dateFormat), nil
}
