package utils

import "time"

const (
	StdTimeFormat = "2006-01-02 15:04:05"
)

// BeginOfWeek t时间戳所在星期，星期一的开始时间。
func BeginOfWeek(t int64) int64 {
	tu := time.Unix(t, 0).UTC()
	offset := int(time.Monday - tu.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(tu.Year(), tu.Month(), tu.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, offset)
	return weekStartDate.Unix()
}

// BeginOfMonth t时间戳所在月份，一号的时间。
func BeginOfMonth(t int64) int64 {
	year, month, _ := time.Unix(t, 0).UTC().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return thisMonth.Unix()
}

// NextMonth t为每月的一号返回下一个月一号
func NextMonth(t int64) int64 {
	tu := time.Unix(t, 0).UTC()
	return tu.AddDate(0, 1, 0).Unix()
}
