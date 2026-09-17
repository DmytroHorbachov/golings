// slices_x008: append к nil
// Make the tests pass!
// I AM NOT DONE
//
// collect собирает положительные числа; nil-срез годится для append.
// Тренирует: нулевое значение среза.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func collect(nums []int) []int {
	var out []int
	for _, n := range nums {
		if n >= 0 {
			out = append(out, n)
		}
	}
	return out
}

func TestCollect(t *testing.T) {
	if got := collect([]int{-1, 0, 2, 5}); !reflect.DeepEqual(got, []int{2, 5}) {
		t.Errorf("collect = %v", got)
	}
}
