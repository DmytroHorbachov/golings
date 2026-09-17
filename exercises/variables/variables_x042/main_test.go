// variables_x042: Неизменяемые строки
// Make the tests pass!
// I AM NOT DONE
//
// Функция capitalize должна сделать первую букву ASCII-строки заглавной.
// Код не компилируется: байт строки нельзя изменить.
// Тренирует: неизменяемость строк и преобразование в []byte.
// Сложность: hard
package main_test

import "testing"

func capitalize(s string) string {
	if s == "" {
		return s
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		s[0] -= 'a' - 'A'
	}
	return s
}

func TestCapitalize(t *testing.T) {
	cases := map[string]string{"gopher": "Gopher", "Go": "Go", "": "", "1st": "1st"}
	for in, want := range cases {
		if got := capitalize(in); got != want {
			t.Errorf("capitalize(%q) = %q, want %q", in, got, want)
		}
	}
}
