// structs_x090: Утверждение к значению или указателю
// Make the tests pass!
// I AM NOT DONE
//
// area получает фигуру через интерфейс; фигуры передаются указателями.
// Утверждение типа к значению Square не срабатывает.
// Тренирует: в интерфейсе хранится *Square, а не Square.
// Сложность: hard
package main_test

import "testing"

type Shape interface{ Name() string }

type Square struct{ Side float64 }

func (s *Square) Name() string { return "square" }

func area(s Shape) float64 {
	if sq, ok := s.(Square); ok {
		return sq.Side * sq.Side
	}
	return 0
}

func TestArea(t *testing.T) {
	if got := area(&Square{3}); got != 9 {
		t.Errorf("area = %v, want 9", got)
	}
}
