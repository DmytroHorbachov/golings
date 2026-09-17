// generics_x009: Обобщённый стек
// Make the tests pass!
// I AM NOT DONE
//
// Stack[T] хранит элементы любого типа; Pop возвращает последний.
// Тренирует: обобщённые типы и методы.
// Сложность: easy
package main_test

import "testing"

type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func (s *Stack[T]) Pop() T {
	v := s.items[0]
	s.items = s.items[:len(s.items)-1]
	return v
}

func TestStack(t *testing.T) {
	var s Stack[string]
	s.Push("a")
	s.Push("b")
	if s.Pop() != "b" || s.Pop() != "a" {
		t.Errorf("Stack works incorrectly")
	}
}
