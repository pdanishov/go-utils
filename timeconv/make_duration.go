package timeconv

import "time"

// Собрать продолжительность по входящим значениям времени.
func MakeDuration(hour, minute, second time.Duration) time.Duration {
	return time.Duration(hour*time.Hour + minute*time.Minute + second*time.Second)
}
