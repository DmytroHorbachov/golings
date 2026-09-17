// structs81
// Make the tests pass!

// I AM NOT DONE
//
// isEmpty проверяет, что у точки нулевые координаты.
// Тренирует: нулевое значение структуры и сравнение ==.
// Сложность: easy
package main_test

import "testing"

type Point struct{ X, Y int }

func isEmpty(p Point) bool {
	return p != Point{}
}

func TestIsEmpty(t *testing.T) {
	var p Point
	if !isEmpty(p) || isEmpty(Point{1, 0}) {
		t.Errorf("isEmpty works incorrectly")
	}
}
