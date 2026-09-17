// structs_x021: Реализация интерфейса
// Make the tests pass!
// I AM NOT DONE
//
// Circle должен реализовать интерфейс Shape с методом Perimeter.
// Тренирует: неявную реализацию интерфейса.
// Сложность: easy
package main_test

import "testing"

type Shape interface{ Perimeter() float64 }

type Circle struct{ R float64 }

func (c Circle) Perimetr() float64 { return 2 * 3 * c.R }

func TestCircleShape(t *testing.T) {
	var s Shape = Circle{R: 2}
	if s.Perimeter() != 12 {
		t.Errorf("Perimeter = %v", s.Perimeter())
	}
}
