// slices87
// Make the tests pass!

// I AM NOT DONE
//
// take возвращает первые n элементов (n не больше длины).
// Тренирует: s[:n].
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func take(s []int, n int) []int {
	return s[n:]
}

func TestTake(t *testing.T) {
	if got := take([]int{1, 2, 3, 4}, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("take = %v", got)
	}
}
