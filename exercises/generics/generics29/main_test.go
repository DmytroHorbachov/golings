// generics29
// Make the tests pass!

// I AM NOT DONE
//
// Box[int] и Box[int64] — разные типы. Код не компилируется при присваивании.
// Тренирует: инстанциации обобщённого типа несовместимы между собой.
// Сложность: hard
package main_test

import "testing"

type Box[T any] struct{ V T }

func widen(b Box[int]) Box[int64] {
	return b
}

func TestWiden(t *testing.T) {
	if widen(Box[int]{7}).V != 7 {
		t.Errorf("widen failed")
	}
}
