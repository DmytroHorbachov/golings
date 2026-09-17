// functions_x059: Попарное объединение
// Make the tests pass!
// I AM NOT DONE
//
// zipWith должна объединить два среза поэлементно с помощью f.
// Длина результата — длина более короткого среза.
// Тренирует: функции высшего порядка над двумя последовательностями.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func zipWith(a, b []int, f func(int, int) int) []int {
	out := make([]int, len(a))
	for i := range a {
		out[i] = f(a[i], b[i])
	}
	return out
}

func TestZipWith(t *testing.T) {
	mul := func(x, y int) int { return x * y }
	if got := zipWith([]int{1, 2, 3}, []int{4, 5}, mul); !reflect.DeepEqual(got, []int{4, 10}) {
		t.Errorf("zipWith = %v, want [4 10]", got)
	}
	if got := zipWith(nil, []int{1}, mul); len(got) != 0 {
		t.Errorf("zipWith(nil, ...) = %v, want empty", got)
	}
}
