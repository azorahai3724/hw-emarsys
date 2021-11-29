package duedatecalc

import (
	"fmt"
	"time"
)

const (
	dateFormat    = "2006-01-02 15:04"
	timeFormat    = "15:04"
	workDayStart  = "09:00"
	workDayEnd    = "17:00"
	workDayLength = 8
	overNight     = 16
	oneDay        = 24
)

func DueDateCalculator(issueStart string, turnAround int) (string, error) {
	t, err := time.Parse(dateFormat, issueStart)
	if err != nil {
		return "", fmt.Errorf("parse issue start: %w", err)
	}

	if isWeekend(t) {
		return "", &InputWeekendErr{}
	}

	b, err := isWithinWorkingHours(workDayStart, workDayEnd, t)

	if err != nil {
		return "", fmt.Errorf("check input is within working hours")
	}
	if !b {
		return "", &NotWithinWorkingHoursErr{}
	}

	if !isValidTurnaround(turnAround) {
		return "", &InvalidTurnaroundErr{}
	}

	newTime := t.Add(time.Hour * 8)

	return newTime.Format(dateFormat), nil
}
