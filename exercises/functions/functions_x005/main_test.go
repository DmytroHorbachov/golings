// functions_x005: Функция как аргумент
// Make the tests pass!
// I AM NOT DONE
//
// Функция transform применяет f к каждому элементу.
// Нужно получить квадраты, но передана не та функция.
// Тренирует: передачу функции как значения.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func double(x int) int { return x * 2 }
func square(x int) int { return x * x }

func transform(nums []int, f func(int) int) []int {
	out := make([]int, len(nums))
	for i, n := range nums {
		out[i] = f(n)
	}
	return out
}

func squares(nums []int) []int {
	return transform(nums, double)
}

func TestSquares(t *testing.T) {
	if got := squares([]int{1, 3, 4}); !reflect.DeepEqual(got, []int{1, 9, 16}) {
		t.Errorf("squares(1,3,4) = %v, want [1 9 16]", got)
	}
}
