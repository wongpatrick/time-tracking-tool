package model

import "time"

type Shift struct {
	ShiftID    int64     `json:"ShiftID"`
	EmployeeID int64     `json:"EmployeeID"`
	StartTime  time.Time `json:"StartTime"`
	EndTime    time.Time `json:"EndTime"`
}

type EmployeeSummary struct {
	EmployeeID    int64   `json:"EmployeeID"`
	StartOfWeek   string  `json:"StartOfWeek"`
	RegularHours  float64 `json:"RegularHours"`
	OvertimeHours float64 `json:"OvertimeHours"`
	InvalidShifts []int64 `json:"InvalidShifts"`
}
