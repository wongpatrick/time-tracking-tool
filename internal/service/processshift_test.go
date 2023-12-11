package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wongpatrick/time-tracking-tool/internal/model"
)

func TestProcessShift(t *testing.T) {
	var test = map[string]struct {
		shifts   []model.Shift
		expected []model.EmployeeSummary
	}{
		"when there are shifts then return employee summary": {
			shifts: []model.Shift{
				{
					ShiftID:    1,
					EmployeeID: 1,
					StartTime:  time.Date(2023, time.January, 1, 8, 0, 0, 0, time.UTC),
					EndTime:    time.Date(2023, time.January, 1, 16, 0, 0, 0, time.UTC),
				},
			},
			expected: []model.EmployeeSummary{
				{
					EmployeeID:    1,
					StartOfWeek:   "2023-01-01",
					RegularHours:  8,
					OvertimeHours: 0,
					InvalidShifts: []int64{},
				},
			},
		},
	}

	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			actual := ProcessShifts(tt.shifts)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
