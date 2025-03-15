// Package jptime provides functions for handling time in Japan.
package jptime

import (
	"fmt"
	"time"
)

// NowJST returns the current time in Japan Standard Time (JST).
func NowJST() (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load location: %w", err)
	}
	return time.Now().In(loc), nil
}

// ToJST converts a time to Japan Standard Time (JST).
func ToJST(t time.Time) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load location: %w", err)
	}
	return t.In(loc), nil
}

// TimeFormatJST formats a time in the Japanese date format (yyyy年MM月dd日).
func TimeFormatJST(t time.Time) string {
	return t.Format("2006年01月02日 15時04分05秒")
}

// WeekdaysJPN returns a slice of the weekdays in Japanese (Sunday to Saturday).
func WeekdaysJPN() []string {
	return []string{"日", "月", "火", "水", "木", "金", "土"}
}

// DayOfWeekJPN returns the Japanese weekday name for the given date.
func DayOfWeekJPN(t time.Time) string {
	weekdays := WeekdaysJPN()
	return weekdays[t.Weekday()]
}
