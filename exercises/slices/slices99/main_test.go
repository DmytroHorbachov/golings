// slices99
// Make the tests pass!

// I AM NOT DONE
//
// lastOr возвращает последний элемент или def для пустого среза.
// Тренирует: индекс последнего элемента.
// Сложность: easy
package main_test

import "testing"

func lastOr(s []int, def int) int {
	if len(s) == 0 {
		return def
	}
	return s[0]
}

func TestLastOr(t *testing.T) {
	if lastOr([]int{4, 5, 6}, 0) != 6 || lastOr(nil, -1) != -1 {
		t.Errorf("lastOr works incorrectly")
	}
}
