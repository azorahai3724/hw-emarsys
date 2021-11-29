package duedatecalc

type InputWeekendErr struct{}

func (i *InputWeekendErr) Error() string {
	return "starting time cannot be on a weekend"
}
