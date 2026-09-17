// anonymous_functions31
// Make the tests pass!

// I AM NOT DONE
//
// normalize применяет по порядку срез литералов-преобразований.
// Тренирует: срез функциональных литералов.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func normalize(s string) string {
	steps := []func(string) string{
		strings.TrimSpace,
		func(x string) string { return strings.ReplaceAll(x, "  ", " ") },
		strings.ToLower,
	}
	for _, f := range steps {
		f(s)
	}
	return s
}

func TestNormalize(t *testing.T) {
	if got := normalize("  Hello  World "); got != "hello world" {
		t.Errorf("normalize = %q", got)
	}
}
