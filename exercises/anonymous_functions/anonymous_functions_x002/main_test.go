// anonymous_functions_x002: Литерал в переменной
// Make the tests pass!
// I AM NOT DONE
//
// square хранится в переменной и вызывается для каждого числа.
// Тренирует: присваивание функционального литерала переменной.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func squares(nums []int) []int {
	square := func(x int) int {
		return x + x
	}
	var out []int
	for _, n := range nums {
		out = append(out, square(n))
	}
	return out
}

func TestSquares(t *testing.T) {
	if got := squares([]int{1, 3, 5}); !reflect.DeepEqual(got, []int{1, 9, 25}) {
		t.Errorf("squares = %v", got)
	}
}
