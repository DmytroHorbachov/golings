// structs_x005: Метод-вычисление
// Make the tests pass!
// I AM NOT DONE
//
// Area возвращает площадь прямоугольника.
// Тренирует: методы со значимым получателем.
// Сложность: easy
package main_test

import "testing"

type Rect struct{ W, H float64 }

func (r Rect) Area() float64 {
	return r.W + r.H
}

func TestArea(t *testing.T) {
	if got := (Rect{3, 4}).Area(); got != 12 {
		t.Errorf("Area = %v", got)
	}
}
