// generics34
// Make the tests pass!

// I AM NOT DONE
//
// SumBy суммирует значения, извлечённые из элементов функцией.
// Тренирует: два параметра типа: элемент и числовой результат.
// Сложность: medium
package main_test

import "testing"

type Number interface{ ~int | ~float64 }

func SumBy[T any, N Number](s []T, f func(T) N) N {
	var total N
	for _, v := range s {
		total = f(v)
	}
	return total
}

type Line struct {
	Price float64
	Qty   int
}

func orderTotal(lines []Line) float64 {
	return SumBy(lines, func(l Line) float64 { return l.Price })
}

func TestOrderTotal(t *testing.T) {
	if got := orderTotal([]Line{{1.5, 2}, {2, 3}}); got != 9 {
		t.Errorf("orderTotal = %v, want 9", got)
	}
}
