// primitive_types39
// Make the tests pass!

// I AM NOT DONE
//
// isRussian проверяет, что строка состоит только из строчных русских букв.
// Слова с буквой «ё» сейчас отклоняются.
// Тренирует: 'ё' находится вне непрерывного диапазона 'а'..'я' в Unicode.
// Сложность: hard
package main_test

import "testing"

func isRussian(s string) bool {
	for _, r := range s {
		if r < 'а' || r > 'я' {
			return false
		}
	}
	return s != ""
}

func TestIsRussian(t *testing.T) {
	cases := map[string]bool{"ёлка": true, "мир": true, "ещё": true, "hello": false, "Мир": false, "": false}
	for in, want := range cases {
		if got := isRussian(in); got != want {
			t.Errorf("isRussian(%q) = %v, want %v", in, got, want)
		}
	}
}
