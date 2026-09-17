// variables30
// Make the tests pass!

// I AM NOT DONE
//
// Функция divider должна вернуть строку из n символов '-'.
// Код не компилируется: строку нельзя умножить на число.
// Тренирует: работа со строковыми значениями через пакет strings.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func divider(n int) string {
	return "-" * n
}

func TestDivider(t *testing.T) {
	if got := divider(5); got != "-----" {
		t.Errorf("divider(5) = %q, want %q", got, "-----")
	}
	if got := divider(0); got != "" {
		t.Errorf("divider(0) = %q, want empty", got)
	}
}
