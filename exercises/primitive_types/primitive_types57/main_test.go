// primitive_types57
// Make the tests pass!

// I AM NOT DONE
//
// censor должна заменить все цифры в строке на '*'.
// Байты меняются, но функция возвращает исходную строку.
// Тренирует: []byte(s) создаёт копию; строка при этом не меняется.
// Сложность: hard
package main_test

import "testing"

func censor(s string) string {
	b := []byte(s)
	for i := range b {
		if b[i] >= '0' && b[i] <= '9' {
			b[i] = '*'
		}
	}
	return s
}

func TestCensor(t *testing.T) {
	if got := censor("card 1234"); got != "card ****" {
		t.Errorf("censor = %q, want %q", got, "card ****")
	}
}
