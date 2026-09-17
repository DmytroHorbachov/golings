// range_x001: Пропуск индекса
// Make the tests pass!
// I AM NOT DONE
//
// sum складывает элементы среза с помощью range.
// Тренирует: пустой идентификатор для ненужного индекса.
// Сложность: easy
package main_test

import "testing"

func sum(nums []int) int {
	total := 0
	for v := range nums {
		total += v
	}
	return total
}

func TestSum(t *testing.T) {
	if got := sum([]int{10, 20, 30}); got != 60 {
		t.Errorf("sum = %d, want 60", got)
	}
}
