// generics59
// Make the tests pass!

// I AM NOT DONE
//
// Pair[K, V] хранит ключ и значение; Swap меняет их местами.
// Тренирует: обобщённую структуру с двумя параметрами.
// Сложность: easy
package main_test

import "testing"

type Pair[K, V any] struct {
	Key K
	Val V
}

func Swap[K, V any](p Pair[K, V]) Pair[V, K] {
	return Pair[V, K]{Key: p.Val}
}

func TestSwap(t *testing.T) {
	s := Swap(Pair[string, int]{"age", 30})
	if s.Key != 30 || s.Val != "age" {
		t.Errorf("Swap = %+v", s)
	}
}
