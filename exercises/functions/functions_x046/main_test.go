// functions_x046: Пробрасывание вариативных аргументов
// Make the tests pass!
// I AM NOT DONE
//
// logLine должна добавить префикс и передать все аргументы дальше в format.
// Сейчас аргументы теряются.
// Тренирует: сборку нового среза аргументов и передачу через ...
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func format(parts ...string) string {
	return strings.Join(parts, " ")
}

func logLine(prefix string, args ...string) string {
	return format(prefix)
}

func TestLogLine(t *testing.T) {
	if got := logLine("[info]", "server", "started"); got != "[info] server started" {
		t.Errorf("logLine = %q", got)
	}
	if got := logLine("[warn]"); got != "[warn]" {
		t.Errorf("logLine without args = %q", got)
	}
}
