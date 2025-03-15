package jptime

import (
	"testing"
	"time"
)

func TestNowJST(t *testing.T) {
	jst, err := NowJST()
	if err != nil {
		t.Errorf("NowJST() failed: %v", err)
	}
	if jst.Location().String() != "Asia/Tokyo" {
		t.Errorf("NowJST() returned %v, expected Asia/Tokyo", jst.Location().String())
	}
}

func TestToJST(t *testing.T) {
	utc := time.Now().UTC()
	jst, err := ToJST(utc)
	if err != nil {
		t.Errorf("ToJST() failed: %v", err)
	}
	if jst.Location().String() != "Asia/Tokyo" {
		t.Errorf("ToJST() returned %v, expected Asia/Tokyo", jst.Location().String())
	}
}

func TestTimeFormatJST(t *testing.T) {
	jst := time.Date(2025, 3, 14, 15, 4, 5, 0, time.FixedZone("JST", 9*60*60))
	formatted := TimeFormatJST(jst)
	if formatted != "2025年03月14日 15時04分05秒" {
		t.Errorf("TimeFormatJST() returned %s, expected 2025年03月14日 15時04分05秒", formatted)
	}
}
func TestWeekdaysJPN(t *testing.T) {
	weekdays := WeekdaysJPN()
	if len(weekdays) != 7 {
		t.Errorf("WeekdaysJPN() returned %d weekdays, expected 7", len(weekdays))
	}
	if weekdays[0] != "日" {
		t.Errorf("WeekdaysJPN() returned %s for Sunday, expected '日'", weekdays[0])
	}
	if weekdays[1] != "月" {
		t.Errorf("WeekdaysJPN() returned %s for Monday, expected '月'", weekdays[1])
	}
	if weekdays[2] != "火" {
		t.Errorf("WeekdaysJPN() returned %s for Tuesday, expected '火'", weekdays[2])
	}
	if weekdays[3] != "水" {
		t.Errorf("WeekdaysJPN() returned %s for Wednesday, expected '水'", weekdays[3])
	}
	if weekdays[4] != "木" {
		t.Errorf("WeekdaysJPN() returned %s for Thursday, expected '木'", weekdays[4])
	}
	if weekdays[5] != "金" {
		t.Errorf("WeekdaysJPN() returned %s for Friday, expected '金'", weekdays[5])
	}
	if weekdays[6] != "土" {
		t.Errorf("WeekdaysJPN() returned %s for Saturday, expected '土'", weekdays[6])
	}
}

func TestDayOfWeekJPN(t *testing.T) {
	cases := []struct {
		date     time.Time
		expected string
	}{
		{time.Date(1900, 1, 1, 0, 0, 0, 0, time.FixedZone("JST", 9*60*60)), "月"},
		{time.Date(2000, 1, 1, 0, 0, 0, 0, time.FixedZone("JST", 9*60*60)), "土"},
		{time.Date(2025, 1, 1, 0, 0, 0, 0, time.FixedZone("JST", 9*60*60)), "水"},
	}
	for _, c := range cases {
		dayOfWeek := DayOfWeekJPN(c.date)
		if dayOfWeek != c.expected {
			t.Errorf("DayOfWeekJPN() returned %s for %s, expected %s", dayOfWeek, c.date, c.expected)
		}
	}
}
