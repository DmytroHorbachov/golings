// variables40
// Make the tests pass!

// I AM NOT DONE
//
// Функция charCount должна вернуть количество символов, а не байт.
// Для "héllo" ожидается 5.
// Тренирует: len(string) возвращает байты; руны считаются отдельно.
// Сложность: medium
package main_test

import (
	"testing"
	"unicode/utf8"
)

func charCount(s string) int {
	_ = utf8.RuneLen
	return len(s)
}

func TestCharCount(t *testing.T) {
	cases := map[string]int{"hello": 5, "héllo": 5, "日本": 2, "": 0}
	for in, want := range cases {
		if got := charCount(in); got != want {
			t.Errorf("charCount(%q) = %d, want %d", in, got, want)
		}
	}
}
