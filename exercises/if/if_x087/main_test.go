// if_x087: Сравнение указателей
// Make the tests pass!
// I AM NOT DONE
//
// samePoint должна сравнивать точки по координатам.
// Для разных, но одинаковых по значению точек возвращается false.
// Тренирует: == на указателях сравнивает адреса, а не значения.
// Сложность: hard
package main_test

import "testing"

type Point struct{ X, Y int }

func samePoint(a, b *Point) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a == b {
		return true
	}
	return false
}

func TestSamePoint(t *testing.T) {
	if !samePoint(&Point{1, 2}, &Point{1, 2}) {
		t.Errorf("equal points should match")
	}
	if samePoint(&Point{1, 2}, &Point{2, 1}) {
		t.Errorf("different points should not match")
	}
	if !samePoint(nil, nil) || samePoint(nil, &Point{}) {
		t.Errorf("nil handling is wrong")
	}
}
