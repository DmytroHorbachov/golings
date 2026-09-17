// primitive_types_x008: Верхний регистр
// Make the tests pass!
// I AM NOT DONE
//
// shout должна перевести строку в верхний регистр.
// Тренирует: функции пакета strings.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func shout(s string) string {
	return strings.ToTitle(strings.ToLower(s)) + "!"
}

func TestShout(t *testing.T) {
	if got := shout("go fast"); got != "GO FAST" {
		t.Errorf("shout = %q, want GO FAST", got)
	}
}
