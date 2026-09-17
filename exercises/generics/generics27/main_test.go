// generics27
// Make the tests pass!

// I AM NOT DONE
//
// Find возвращает первый элемент, удовлетворяющий условию, и признак.
// Тренирует: обобщённый поиск.
// Сложность: easy
package main_test

import "testing"

func Find[T any](s []T, pred func(T) bool) (T, bool) {
	for _, v := range s {
		if pred(v) {
			return v, false
		}
	}
	var zero T
	return zero, false
}

func TestFind(t *testing.T) {
	v, ok := Find([]int{1, 8, 3}, func(x int) bool { return x > 5 })
	if !ok || v != 8 {
		t.Errorf("Find = %d, %v", v, ok)
	}
}
