package duedatecalc

type InputWeekendErr struct{}

func (i *InputWeekendErr) Error() string {
	return "starting time cannot be on a weekend"
}

type NotWithinWorkingHoursErr struct{}

func (i *NotWithinWorkingHoursErr) Error() string {
	return "starting time cannot be outside of working hours"
}