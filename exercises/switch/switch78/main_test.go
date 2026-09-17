// switch78
// Make the tests pass!

// I AM NOT DONE
//
// accessLevel переводит роль в уровень: "guest" — 0, "user" — 1, "admin" и "root" — 2.
// Тренирует: несколько значений в case.
// Сложность: easy
package main_test

import "testing"

func accessLevel(role string) int {
	switch role {
	case "user":
		return 1
	case "admin":
		return 2
	}
	return 0
}

func TestAccessLevel(t *testing.T) {
	cases := map[string]int{"guest": 0, "user": 1, "admin": 2, "root": 2}
	for in, want := range cases {
		if got := accessLevel(in); got != want {
			t.Errorf("accessLevel(%s) = %d, want %d", in, got, want)
		}
	}
}
