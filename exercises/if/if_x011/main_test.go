// if_x011: Минимальная длина пароля
// Make the tests pass!
// I AM NOT DONE
//
// strongEnough должна требовать пароль длиной не менее 8 символов.
// Тренирует: условие на длину строки.
// Сложность: easy
package main_test

import "testing"

func strongEnough(pw string) bool {
	if len(pw) < 9 {
		return false
	}
	return true
}

func TestStrongEnough(t *testing.T) {
	cases := map[string]bool{"1234567": false, "12345678": true, "correct horse": true}
	for in, want := range cases {
		if got := strongEnough(in); got != want {
			t.Errorf("strongEnough(%q) = %v, want %v", in, got, want)
		}
	}
}
