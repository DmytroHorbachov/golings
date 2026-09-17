// if_x041: Простая проверка email
// Make the tests pass!
// I AM NOT DONE
//
// validEmail: ровно один символ @, он не первый, и после него есть точка
// (не сразу после @ и не в конце).
// Тренирует: последовательность проверок с ранним возвратом.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func validEmail(s string) bool {
	at := strings.Index(s, "@")
	if at < 0 {
		return false
	}
	return strings.Contains(s, ".")
}

func TestValidEmail(t *testing.T) {
	cases := map[string]bool{
		"ann@mail.com": true, "a@b.co": true, "@mail.com": false, "ann.mail.com": false,
		"ann@mail": false, "ann@.com": false, "ann@mail.": false, "a@b@c.com": false,
	}
	for in, want := range cases {
		if got := validEmail(in); got != want {
			t.Errorf("validEmail(%q) = %v, want %v", in, got, want)
		}
	}
}
