// generics_x001: Параметр типа
// Make the tests pass!
// I AM NOT DONE
//
// Map применяет функцию к каждому элементу среза любого типа.
// Код не компилируется: параметр типа объявлен без ограничения.
// Тренирует: синтаксис [T any].
// Сложность: easy
package main_test

import (
	"reflect"
	"strconv"
	"testing"
)

func Map[T, U](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

func TestMap(t *testing.T) {
	if got := Map([]int{1, 2}, strconv.Itoa); !reflect.DeepEqual(got, []string{"1", "2"}) {
		t.Errorf("Map = %v", got)
	}
}
