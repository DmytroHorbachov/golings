// anonymous_functions_x024: Посетитель списка
// Make the tests pass!
// I AM NOT DONE
//
// each вызывает колбэк для каждого значения; sum суммирует их литералом.
// Тренирует: литерал, изменяющий захваченную переменную.
// Сложность: easy
package main_test

import "testing"

func each(s []int, f func(int)) {
	for _, v := range s {
		f(v)
	}
}

func sum(s []int) int {
	total := 0
	each(s, func(v int) { total = v })
	return total
}

func TestSum(t *testing.T) {
	if sum([]int{1, 2, 3}) != 6 {
		t.Errorf("sum = %d", sum([]int{1, 2, 3}))
	}
}
