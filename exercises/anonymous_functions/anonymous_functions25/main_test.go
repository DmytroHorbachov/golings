// anonymous_functions25
// Make the tests pass!

// I AM NOT DONE
//
// compute должна вернуть значение, увеличенное отложенным литералом.
// Литерал меняет локальную переменную, но результат уже вычислен.
// Тренирует: return с выражением фиксирует значение до выполнения defer.
// Сложность: hard
package main_test

import "testing"

func compute(x int) int {
	result := x * 2
	defer func() { result++ }()
	return result
}

func TestCompute(t *testing.T) {
	if got := compute(5); got != 11 {
		t.Errorf("compute(5) = %d, want 11", got)
	}
}
