package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wongpatrick/time-tracking-tool/internal/model"
)

func TestCalculateEmployeeSummary(t *testing.T) {
	var test = map[string]struct {
		employeeID    int64
		shifts        []model.Shift
		invalidShifts map[int64]bool
		expected      []model.EmployeeSummary
	}{
		"when Valid EmployeeID and Shifts then return EmployeeSummary": {
			employeeID: 1,
			shifts: []model.Shift{
				{
					ShiftID:   1,
					StartTime: time.Date(2023, time.January, 15, 8, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2023, time.January, 15, 12, 0, 0, 0, time.UTC),
				},
			},
			invalidShifts: map[int64]bool{},
			expected: []model.EmployeeSummary{{
				EmployeeID:    1,
				StartOfWeek:   "2023-01-15",
				RegularHours:  4.0,
				OvertimeHours: 0,
				InvalidShifts: []int64{},
			}},
		},
	}

	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			actual := calculateEmployeeSummary(tt.employeeID, tt.shifts, tt.invalidShifts)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetStartOfWeek(t *testing.T) {
	var test = map[string]struct {
		inputTime time.Time
		expected  time.Time
	}{
		"when Sunday then return Sunday": {
			inputTime: time.Date(2023, time.January, 15, 12, 0, 0, 0, time.UTC),
			expected:  time.Date(2023, time.January, 15, 12, 0, 0, 0, time.UTC),
		},
		"when Wednesday then return Sunday": {
			inputTime: time.Date(2023, time.January, 18, 18, 0, 0, 0, time.UTC),
			expected:  time.Date(2023, time.January, 15, 18, 0, 0, 0, time.UTC),
		},
	}

	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			actual := getStartOfWeek(tt.inputTime)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
