// functions_x086: Копия strings.Builder
// Make the tests pass!
// I AM NOT DONE
//
// render собирает текст, передавая strings.Builder во вспомогательную функцию.
// Код паникует: «illegal use of non-zero Builder copied by value».
// Тренирует: некоторые типы нельзя копировать после начала использования.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func writeLine(b strings.Builder, s string) {
	b.WriteString(s)
	b.WriteString("\n")
}

func render(lines []string) string {
	var b strings.Builder
	b.WriteString("BEGIN\n")
	for _, l := range lines {
		writeLine(b, l)
	}
	b.WriteString("END")
	return b.String()
}

func TestRender(t *testing.T) {
	want := "BEGIN\nx\ny\nEND"
	if got := render([]string{"x", "y"}); got != want {
		t.Errorf("render = %q, want %q", got, want)
	}
}
