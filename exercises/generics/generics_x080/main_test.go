// generics_x080: comparable и интерфейсы
// Make the tests pass!
// I AM NOT DONE
//
// CountEqual считает элементы, равные x. При инстанциации типом any
// сравнение срезов внутри интерфейсов паникует во время выполнения.
// Тренирует: comparable допускает интерфейсные типы, но == может паниковать.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func CountEqual[T comparable](s []T, x T) int {
	n := 0
	for _, v := range s {
		if v == x {
			n++
		}
	}
	return n
}

func TestCountEqual(t *testing.T) {
	_ = reflect.DeepEqual
	vals := []any{1, []int{1}, "a", []int{1}}
	if got := CountEqual(vals, any([]int{1})); got != 2 {
		t.Errorf("CountEqual = %d, want 2", got)
	}
}
