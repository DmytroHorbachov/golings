// switch_x046: Экранирование
// Make the tests pass!
// I AM NOT DONE
//
// escape заменяет перевод строки на \n, табуляцию на \t и обратный слеш на \\.
// Тренирует: switch по руне внутри цикла со strings.Builder.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func escape(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		case '\n':
			b.WriteString(`\n`)
		case '\t':
			b.WriteString(`\n`)
		}
	}
	return b.String()
}

func TestEscape(t *testing.T) {
	in := "a\tb\nc\\d"
	want := `a\tb\nc\\d`
	if got := escape(in); got != want {
		t.Errorf("escape(%q) = %s, want %s", in, got, want)
	}
}
