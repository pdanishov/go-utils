package timeconv

import "time"

// Форматировать продолжительность по заданному шаблону из пакета time.
func FormatDuration(duration time.Duration, layout string) string {
	return time.Unix(0, 0).UTC().Add(duration).Format(layout)
}
