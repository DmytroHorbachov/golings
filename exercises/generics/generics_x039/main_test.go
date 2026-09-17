// generics_x039: Разбиение на части
// Make the tests pass!
// I AM NOT DONE
//
// Chunk делит срез любого типа на части размера n.
// Тренирует: обобщённые срезы срезов.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func Chunk[T any](s []T, n int) [][]T {
	var out [][]T
	for i := 0; i < len(s); i += n {
		out = append(out, s[i:i+n])
	}
	return out
}

func TestChunk(t *testing.T) {
	got := Chunk([]string{"a", "b", "c"}, 2)
	if !reflect.DeepEqual(got, [][]string{{"a", "b"}, {"c"}}) {
		t.Errorf("Chunk = %v", got)
	}
}
