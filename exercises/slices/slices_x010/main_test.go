// slices_x010: Сумма
// Make the tests pass!
// I AM NOT DONE
//
// total складывает все элементы среза.
// Тренирует: range по срезу.
// Сложность: easy
package main_test

import "testing"

func total(s []int) int {
	sum := 0
	for _, v := range s {
		sum = v
	}
	return sum
}

func TestTotal(t *testing.T) {
	if got := total([]int{5, 10, 15}); got != 30 {
		t.Errorf("total = %d, want 30", got)
	}
}
