// generics52
// Make the tests pass!

// I AM NOT DONE
//
// Set[T] позволяет добавлять элементы и проверять наличие.
// Тренирует: обобщённый тип на основе map.
// Сложность: easy
package main_test

import "testing"

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) { s[v] = struct{}{} }

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return !ok
}

func TestSet(t *testing.T) {
	s := Set[int]{}
	s.Add(3)
	if !s.Has(3) || s.Has(4) {
		t.Errorf("Set works incorrectly")
	}
}
