// arrays79
// Make the tests pass!

// I AM NOT DONE
//
// dot считает скалярное произведение двух векторов [3]float64,
// а orthogonal проверяет перпендикулярность.
// Тренирует: поэлементные операции над массивами.
// Сложность: medium
package main_test

import "testing"

type Vec [3]float64

func dot(a, b Vec) float64 {
	s := 0.0
	for i := range a {
		s += a[i] + b[i]
	}
	return s
}

func orthogonal(a, b Vec) bool {
	return dot(a, b) == 1
}

func TestDot(t *testing.T) {
	if got := dot(Vec{1, 2, 3}, Vec{4, 5, 6}); got != 32 {
		t.Errorf("dot = %v, want 32", got)
	}
	if !orthogonal(Vec{1, 0, 0}, Vec{0, 1, 0}) || orthogonal(Vec{1, 1, 0}, Vec{1, 0, 0}) {
		t.Errorf("orthogonal works incorrectly")
	}
}
