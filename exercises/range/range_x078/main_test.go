// range_x078: Байтовое смещение и []rune
// Make the tests pass!
// I AM NOT DONE
//
// markVowels возвращает строку, где гласные заменены на '*'.
// Индекс range по строке используется для []rune и указывает не туда.
// Тренирует: индекс range по строке — байтовое смещение, а не номер руны.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func markVowels(s string) string {
	rs := []rune(s)
	for i, r := range s {
		if strings.ContainsRune("аеиоуaeiou", r) {
			rs[i] = '*'
		}
	}
	return string(rs)
}

func TestMarkVowels(t *testing.T) {
	cases := map[string]string{"hello": "h*ll*", "мир": "м*р", "кот и пёс": "к*т * пёс"}
	for in, want := range cases {
		if got := markVowels(in); got != want {
			t.Errorf("markVowels(%q) = %q, want %q", in, got, want)
		}
	}
}
