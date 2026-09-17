// if_x069: Область видимости if-init
// Make the tests pass!
// I AM NOT DONE
//
// parseOr должна вернуть разобранное число или def.
// Код не компилируется: переменная из if используется после блока.
// Тренирует: переменные из инициализатора if видны только внутри if/else.
// Сложность: hard
package main_test

import (
	"strconv"
	"testing"
)

func parseOr(s string, def int) int {
	if n, err := strconv.Atoi(s); err != nil {
		return def
	}
	return n
}

func TestParseOr(t *testing.T) {
	if got := parseOr("12", 0); got != 12 {
		t.Errorf("parseOr(12) = %d", got)
	}
	if got := parseOr("x", 5); got != 5 {
		t.Errorf("parseOr(x, 5) = %d", got)
	}
}
