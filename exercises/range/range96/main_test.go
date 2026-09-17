// range96
// Make the tests pass!

// I AM NOT DONE
//
// product перемножает элементы среза.
// Тренирует: аккумулятор в range.
// Сложность: easy
package main_test

import "testing"

func product(s []int) int {
	p := 0
	for _, v := range s {
		p *= v
	}
	return p
}

func TestProduct(t *testing.T) {
	if got := product([]int{2, 3, 4}); got != 24 {
		t.Errorf("product = %d, want 24", got)
	}
}
