// arrays72
// Make the tests pass!

// I AM NOT DONE
//
// count должна посчитать, сколько раз x встречается в массиве.
// Тренирует: сравнение элементов в цикле.
// Сложность: easy
package main_test

import "testing"

func count(a [6]int, x int) int {
	n := 0
	for _, v := range a {
		if v != x {
			n++
		}
	}
	return n
}

func TestCount(t *testing.T) {
	if got := count([6]int{1, 2, 1, 3, 1, 1}, 1); got != 4 {
		t.Errorf("count = %d, want 4", got)
	}
}
