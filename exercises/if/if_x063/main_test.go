// if_x063: Сравнение с нормализацией
// Make the tests pass!
// I AM NOT DONE
//
// sameTitle сравнивает заголовки без учёта регистра и лишних пробелов по краям.
// Пустые (после обрезки) заголовки никогда не считаются одинаковыми.
// Тренирует: подготовку данных перед условием.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func sameTitle(a, b string) bool {
	if strings.ToLower(a) == strings.ToLower(b) {
		return true
	}
	return false
}

func TestSameTitle(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{{"Go Tour", "  go tour ", true}, {"Go", "Rust", false}, {"  ", "", false}, {"", "", false}}
	for _, c := range cases {
		if got := sameTitle(c.a, c.b); got != c.want {
			t.Errorf("sameTitle(%q, %q) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
