// generics_x071: nil для T
// Make the tests pass!
// I AM NOT DONE
//
// FirstOr возвращает первый элемент или значение «ничего». Код не компилируется:
// nil нельзя использовать как значение параметра типа.
// Тренирует: нулевое значение T получается через var zero T.
// Сложность: hard
package main_test

import "testing"

func FirstOr[T any](s []T) (T, bool) {
	if len(s) == 0 {
		return nil, false
	}
	return s[0], true
}

func TestFirstOr(t *testing.T) {
	if v, ok := FirstOr([]int{}); ok || v != 0 {
		t.Errorf("FirstOr(empty) = %v, %v", v, ok)
	}
	if v, ok := FirstOr([]string{"x"}); !ok || v != "x" {
		t.Errorf("FirstOr = %v, %v", v, ok)
	}
}
