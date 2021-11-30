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

	isWorkingHour, err := isWithinWorkingHours(workDayStart, workDayEnd, t)
	if err != nil {
		return "", fmt.Errorf("check input is within working hours: %w", err)
	}
	if !isWorkingHour {
		return "", &NotWithinWorkingHoursErr{}
	}

	if !isValidTurnaround(turnAround) {
		return "", &InvalidTurnaroundErr{}
	}

	t, err = fastForward(t, turnAround)

	if err != nil {
		return "", fmt.Errorf("fast forward date: %w", err)
	}

	return t.Format(dateFormat), nil
}

func fastForward(t time.Time, turnAround int) (time.Time, error) {
	t, err := fastForwardHours(t, turnAround)
	if err != nil {
		return time.Time{}, fmt.Errorf("fast forward hours: %w", err)
	}

	t = fastForwardDays(t, turnAround)

	return t, nil
}

func fastForwardHours(t time.Time, turnAround int) (time.Time, error) {
	hoursLeft := turnAround % workDayLength

	for {
		if hoursLeft > 0 {
			isWorkingHour, err := isWithinWorkingHours(workDayStart, workDayEnd, t)
			if err != nil {
				return time.Time{}, fmt.Errorf("check input is within working hours: %w", err)
			}
			if isWorkingHour {
				hoursLeft -= 1
			}

			t = t.Add(time.Hour * 1)
		} else {
			return t, nil
		}
	}
}

func fastForwardDays(t time.Time, turnAround int) time.Time {
	daysLeft := turnAround / workDayLength

	for {
		if daysLeft > 0 {
			if !isWeekend(t) {
				daysLeft -= 1
			}
			t = t.AddDate(0, 0, 1)
		} else {
			return t
		}
	}
}
