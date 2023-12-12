package service

import (
	"time"

	"github.com/wongpatrick/time-tracking-tool/internal/model"
)

func ProcessShifts(shifts []model.Shift) []model.EmployeeSummary {
	employeeMap := make(map[int64][]model.Shift)

	// Group shifts by employee ID
	for _, shift := range shifts {
		// Check if shift crosses midnight of Sunday
		if shift.EndTime.After(getStartOfWeek(shift.StartTime).AddDate(0, 0, 7)) {
			// Splits shift into two shifts, one for Saturday and one for Sunday
			employeeMap[shift.EmployeeID] = append(employeeMap[shift.EmployeeID], model.Shift{
				ShiftID:    shift.ShiftID,
				EmployeeID: shift.EmployeeID,
				StartTime:  shift.StartTime,
				EndTime:    time.Date(shift.StartTime.Year(), shift.StartTime.Month(), shift.StartTime.Day()+1, 0, 0, 0, 0, shift.StartTime.Location()),
			})

			employeeMap[shift.EmployeeID] = append(employeeMap[shift.EmployeeID], model.Shift{
				ShiftID:    shift.ShiftID,
				EmployeeID: shift.EmployeeID,
				StartTime:  time.Date(shift.StartTime.Year(), shift.StartTime.Month(), shift.StartTime.Day()+1, 0, 0, 0, 0, shift.StartTime.Location()),
				EndTime:    shift.EndTime,
			})
		} else {
			employeeMap[shift.EmployeeID] = append(employeeMap[shift.EmployeeID], shift)
		}
	}

	var employeeSummaries []model.EmployeeSummary

	for employeeID, employeeShifts := range employeeMap {
		invalidShifts := findInvalidShifts(employeeShifts)
		summary := calculateEmployeeSummary(employeeID, employeeShifts, invalidShifts)
		employeeSummaries = append(employeeSummaries, summary...)
	}

	return employeeSummaries
}
