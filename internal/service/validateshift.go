package service

import "github.com/wongpatrick/time-tracking-tool/internal/model"

func overlaps(shift1, shift2 model.Shift) bool {
	return shift1.EmployeeID == shift2.EmployeeID &&
		(shift1.StartTime.Before(shift2.EndTime) &&
			shift1.EndTime.After(shift2.StartTime))
}

func findInvalidShifts(shifts []model.Shift) map[int64]bool {
	invalidShifts := make(map[int64]bool, 0)

	for i, shift := range shifts {
		if invalidShifts[shift.ShiftID] {
			continue
		}

		for j := i + 1; j < len(shifts); j++ {
			if overlaps(shift, shifts[j]) {
				invalidShifts[shift.ShiftID] = true
				invalidShifts[shifts[j].ShiftID] = true
			}
		}
	}

	return invalidShifts
}
