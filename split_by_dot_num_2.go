package utils

import (
	"errors"
	"fmt"
	"unicode"
	"unicode/utf8"
)

// Разбить строку с двузначным числом на строку из чисел разделённых точкой.
// Если второго знака нет, он заменяется нулём. Если входящая строка пустая,
// возвращается 0.0. Пример: "47" -> "4.7", "2" -> "2.0", "" -> "0.0"
func SplitByDotNum2(nn string) (string, error) {
	if utf8.RuneCountInString(nn) > 2 { // Проверить что длина строки не больше 2.
		return "", errors.New("input string must be of length 2")
	}
	var s = [2]string{"0", "0"}
	// Проверить что оба символа входящей строки являются числами.
	for i, char := range nn {
		if !unicode.IsDigit(char) {
			return "", fmt.Errorf("%s is not a digit", string(char))
		}
		s[i] = string(char)
	}
	return fmt.Sprintf("%s.%s", s[0], s[1]), nil
}
