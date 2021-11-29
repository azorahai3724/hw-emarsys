package duedatecalc

import (
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04"

func DueDateCalculator(issueStart string, turnAround int) (string, error) {
	t, err := time.Parse(timeFormat, issueStart)
	if err != nil {
		return "", fmt.Errorf("parse issue start: %w", err)
	}
	newTime := t.Add(time.Hour * 8)
	return newTime.Format(timeFormat), nil
}
