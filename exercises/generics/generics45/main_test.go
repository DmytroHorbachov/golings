// generics45
// Make the tests pass!

// I AM NOT DONE
//
// Sum складывает элементы среза чисел любого числового типа.
// Тренирует: ограничение Number и нулевое значение T.
// Сложность: easy
package main_test

import "testing"

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total = v
	}
	return total
}

func TestSum(t *testing.T) {
	if Sum([]int{1, 2, 3}) != 6 || Sum([]float64{0.5, 0.25}) != 0.75 {
		t.Errorf("Sum works incorrectly")
	}
}
