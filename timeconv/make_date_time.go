package timeconv

import "time"

// Собрать дату и время по входящим значениям.
// Возвращается объект времени с текущим системным смещением временной зоны.
func MakeDateTime(year int, month time.Month, day int, hour int, minute int, second int) time.Time {
	return time.Date(year, month, day, hour, minute, second, 0, time.Local)
}
