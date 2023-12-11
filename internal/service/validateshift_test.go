package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wongpatrick/time-tracking-tool/internal/model"
)

func TestOverlaps(t *testing.T) {
	var test = map[string]struct {
		firstShift  model.Shift
		secondShift model.Shift
		expected    bool
	}{
		"when 2 valid shift then return false": {
			firstShift: model.Shift{
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 8, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 16, 0, 0, 0, time.UTC),
			},
			secondShift: model.Shift{
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 16, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 23, 0, 0, 0, time.UTC),
			},
			expected: false,
		},
		"when shift 2 overlaps shift 1 then return true return true": {
			firstShift: model.Shift{
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 8, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 16, 0, 0, 0, time.UTC),
			},
			secondShift: model.Shift{
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 9, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 17, 0, 0, 0, time.UTC),
			},
			expected: true,
		},
		"when shift 1 overlaps shift 2 then return true return true": {
			firstShift: model.Shift{
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 8, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 16, 0, 0, 0, time.UTC),
			},
			secondShift: model.Shift{
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 7, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC),
			},
			expected: true,
		},
		"when shift 1 is inside shift 2 then return true return true": {
			firstShift: model.Shift{
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 9, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC),
			},
			secondShift: model.Shift{
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 8, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 16, 0, 0, 0, time.UTC),
			},
			expected: true,
		},
		"when shift 2 is inside shift 1 then return true return true": {
			firstShift: model.Shift{
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 8, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 16, 0, 0, 0, time.UTC),
			},
			secondShift: model.Shift{
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 9, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC),
			},
			expected: true,
		},
	}

	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			actual := overlaps(tt.firstShift, tt.secondShift)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFindInvalidShifts(t *testing.T) {
	var test = map[string]struct {
		shifts   []model.Shift
		expected map[int64]bool
	}{
		"when there are no invalid shifts then return empty map": {
			shifts: []model.Shift{{
				ShiftID:    1,
				EmployeeID: 1,
				StartTime:  time.Date(2023, time.January, 1, 8, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2023, time.January, 1, 16, 0, 0, 0, time.UTC),
			}},
			expected: map[int64]bool{},
		},
		"when there are invalid shifts then return map of invalid shifts": {
			shifts: []model.Shift{
				{
					ShiftID:    1,
					EmployeeID: 1,
					StartTime:  time.Date(2023, time.January, 1, 8, 0, 0, 0, time.UTC),
					EndTime:    time.Date(2023, time.January, 1, 16, 0, 0, 0, time.UTC),
				},
				{
					ShiftID:    2,
					EmployeeID: 1,
					StartTime:  time.Date(2023, time.January, 1, 7, 0, 0, 0, time.UTC),
					EndTime:    time.Date(2023, time.January, 1, 15, 0, 0, 0, time.UTC),
				},
			},
			expected: map[int64]bool{1: true, 2: true},
		},
	}

	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			actual := findInvalidShifts(tt.shifts)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestIsShiftInvalid(t *testing.T) {
	var test = map[string]struct {
		shiftID       int64
		invalidShifts []int64
		expected      bool
	}{
		"when there are no invalid shifts then return false": {
			shiftID:       1,
			invalidShifts: []int64{},
			expected:      false,
		},
		"when there are no matching shiftID and invalid shifts then return false": {
			shiftID:       1,
			invalidShifts: []int64{2},
			expected:      false,
		},
		"when there is a matching shiftID and invalid shifts then return true": {
			shiftID:       1,
			invalidShifts: []int64{1},
			expected:      true,
		},
	}

	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			actual := isShiftInvalid(tt.shiftID, tt.invalidShifts)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
