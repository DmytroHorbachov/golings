// primitive_types_x003: Глагол %d
// Make the tests pass!
// I AM NOT DONE
//
// label должна вернуть строку "item-7" для числа 7.
// Тренирует: форматирование целых чисел через fmt.Sprintf.
// Сложность: easy
package main_test

import (
	"fmt"
	"testing"
)

func label(n int) string {
	return fmt.Sprintf("item-%c", n)
}

func TestLabel(t *testing.T) {
	if got := label(7); got != "item-7" {
		t.Errorf("label(7) = %q, want item-7", got)
	}
}
