// generics12
// Make the tests pass!

// I AM NOT DONE
//
// Zip объединяет два среза разных типов в срез пар; длина — по короткому.
// Тренирует: обобщённая структура как результат.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type Pair[A, B any] struct {
	First  A
	Second B
}

func Zip[A, B any](a []A, b []B) []Pair[A, B] {
	out := make([]Pair[A, B], len(a))
	for i := range a {
		out[i] = Pair[A, B]{a[i], b[i]}
	}
	return out
}

func TestZip(t *testing.T) {
	got := Zip([]string{"a", "b", "c"}, []int{1, 2})
	if !reflect.DeepEqual(got, []Pair[string, int]{{"a", 1}, {"b", 2}}) {
		t.Errorf("Zip = %v", got)
	}
}
