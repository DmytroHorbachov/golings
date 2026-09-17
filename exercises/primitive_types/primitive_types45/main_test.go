// primitive_types45
// Make the tests pass!

// I AM NOT DONE
//
// upper переводит строчные русские буквы в заглавные вычитанием 32.
// Для буквы «ё» результат неверный.
// Тренирует: смещение между регистрами не одинаково для всех символов.
// Сложность: hard
package main_test

import (
	"testing"
	"unicode"
)

func upper(s string) string {
	r := []rune(s)
	for i := range r {
		if r[i] >= 'а' && r[i] <= 'я' || r[i] == 'ё' {
			r[i] -= 32
		}
	}
	return string(r)
}

func TestUpper(t *testing.T) {
	_ = unicode.ToUpper
	cases := map[string]string{"мир": "МИР", "ёж": "ЁЖ", "ещё": "ЕЩЁ"}
	for in, want := range cases {
		if got := upper(in); got != want {
			t.Errorf("upper(%q) = %q, want %q", in, got, want)
		}
	}
}
