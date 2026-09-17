// arrays51
// Make the tests pass!

// I AM NOT DONE
//
// last должна вернуть последний элемент массива.
// Тренирует: индексацию с нуля.
// Сложность: easy
package main_test

import "testing"

func last(a [4]string) string {
	return a[len(a)-2]
}

func TestLast(t *testing.T) {
	if got := last([4]string{"a", "b", "c", "d"}); got != "d" {
		t.Errorf("last = %q, want d", got)
	}
}
