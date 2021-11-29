package duedatecalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDueDateCalculator(t *testing.T) {
	tests := map[string]struct {
		issueStart     string
		turnAroundTime int
		wantDueDate    string
		wantErr        error
	}{
		"within the same workday, 8 hours turnaround time": {issueStart: "2021-11-29 09:00", turnAroundTime: 8, wantDueDate: "2021-11-29 17:00", wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotDueDate, gotErr := DueDateCalculator(tc.issueStart, tc.turnAroundTime)
			if gotErr != nil {
				assert.ErrorIsf(t, gotErr, tc.wantErr, "with issue start date %s and turnaround time %d, got %w want %w", tc.issueStart, tc.turnAroundTime, gotErr, tc.wantErr)
			}
			if gotDueDate != tc.wantDueDate {
				t.Errorf("with issue start date %s and turnaround time %d, got %s want %s", tc.issueStart, tc.turnAroundTime, gotDueDate, tc.wantDueDate)
			}
		})
	}
}
