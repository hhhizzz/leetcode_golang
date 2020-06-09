package _1185

import "time"

func dayOfTheWeek(day int, month int, year int) string {
	d := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return days[d.Weekday()]
}
