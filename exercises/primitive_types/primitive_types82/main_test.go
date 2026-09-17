// primitive_types82
// Make the tests pass!

// I AM NOT DONE
//
// stripPrefix убирает префикс "abc-" из идентификатора ровно один раз.
// Для "abc-cab-1" сейчас удаляется слишком много.
// Тренирует: TrimLeft принимает набор символов (cutset), а не подстроку.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func stripPrefix(id string) string {
	return strings.TrimLeft(id, "abc-")
}

func TestStripPrefix(t *testing.T) {
	cases := map[string]string{"abc-cab-1": "cab-1", "abc-42": "42", "bca-7": "bca-7"}
	for in, want := range cases {
		if got := stripPrefix(in); got != want {
			t.Errorf("stripPrefix(%q) = %q, want %q", in, got, want)
		}
	}
}
