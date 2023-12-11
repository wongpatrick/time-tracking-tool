package service

import (
	"log"
	"sort"
	"time"

	"github.com/wongpatrick/time-tracking-tool/internal/model"
)

func calculateEmployeeSummary(employeeID int64, shifts []model.Shift, invalidShifts map[int64]bool) []model.EmployeeSummary {
	const hoursInWeek = 24 * 7

	summary := []model.EmployeeSummary{}
	j := -1

	sort.Slice(shifts[:], func(i, j int) bool {
		return shifts[i].StartTime.Before(shifts[j].StartTime)
	})

	for i, shift := range shifts {
		var (
			startOfWeekTime time.Time
			err             error
		)
		if len(summary) > 0 {
			startOfWeekTime, err = time.Parse("2006-01-02", summary[j].StartOfWeek)
			if err != nil {
				log.Fatal(err)
			}
		}

		if i == 0 || shift.StartTime.After(startOfWeekTime.AddDate(0, 0, 7)) {
			j++
			summary = append(summary, model.EmployeeSummary{
				EmployeeID:    employeeID,
				StartOfWeek:   getStartOfWeek(shift.StartTime).Format("2006-01-02"),
				RegularHours:  0,
				OvertimeHours: 0,
				InvalidShifts: []int64{},
			})
		}

		if invalidShifts[shift.ShiftID] {
			summary[j].InvalidShifts = append(summary[j].InvalidShifts, shift.ShiftID)
		} else {
			summary[j].RegularHours += shift.EndTime.Sub(shift.StartTime).Hours()
			if summary[j].RegularHours > 40.0 {
				summary[j].OvertimeHours += summary[j].RegularHours - 40.0
				summary[j].RegularHours = 40.0
			}
		}
	}

	return summary
}

func getStartOfWeek(t time.Time) time.Time {
	_, offset := t.Zone()

	return t.Add(-time.Duration(t.Weekday())*24*time.Hour - time.Duration(offset)*time.Second)
}
