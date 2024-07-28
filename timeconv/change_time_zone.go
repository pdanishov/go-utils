package timeconv

import (
	"math"
	"time"
)

// Изменить значение смещения временной зоны не меняя время.
func ChangeTimeZone(timeDate time.Time, timeZoneOffset time.Duration) time.Time {
	var offset = int(math.Round(timeZoneOffset.Seconds()))
	return time.Date(
		timeDate.Year(),
		timeDate.Month(),
		timeDate.Day(),
		timeDate.Hour(),
		timeDate.Minute(),
		timeDate.Second(),
		timeDate.Nanosecond(),
		time.FixedZone("", offset),
	)
}
