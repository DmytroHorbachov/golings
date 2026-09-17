// slices61
// Make the tests pass!

// I AM NOT DONE
//
// clone копирует данные, но возвращает пустой срез.
// Тренирует: copy копирует min(len(dst), len(src)), ёмкость не учитывается.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func clone(s []int) []int {
	out := make([]int, 0, len(s))
	copy(out, s)
	return out
}

func TestClone(t *testing.T) {
	src := []int{1, 2, 3}
	c := clone(src)
	src[0] = 9
	if !reflect.DeepEqual(c, []int{1, 2, 3}) {
		t.Errorf("clone = %v", c)
	}
}
