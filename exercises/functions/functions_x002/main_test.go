// functions_x002: Распаковка среза
// Make the tests pass!
// I AM NOT DONE
//
// Функция sum принимает переменное число аргументов.
// Нужно передать в неё весь срез целиком. Код не компилируется.
// Тренирует: передачу среза в вариативную функцию.
// Сложность: easy
package main_test

import "testing"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func sumSlice(values []int) int {
	return sum(values)
}

func TestSumSlice(t *testing.T) {
	if got := sumSlice([]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("sumSlice(1,2,3,4) = %d, want 10", got)
	}
}
