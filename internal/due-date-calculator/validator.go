package duedatecalc

import "time"

func isWeekend(t time.Time) bool {
	switch t.Weekday() {
	case time.Friday:
		h, _, _ := t.Clock()
		if h >= 12+5 {
			return true
		}
	case time.Saturday:
		return true
	case time.Sunday:
		return true
	}
	return false
}
