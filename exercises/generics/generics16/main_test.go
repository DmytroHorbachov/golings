// generics16
// Make the tests pass!

// I AM NOT DONE
//
// Half делит значение пополам, умножая на 0.5. Код не компилируется:
// в наборе типов есть целые, для которых 0.5 не представимо.
// Тренирует: константы проверяются для всех типов ограничения.
// Сложность: hard
package main_test

import "testing"

type Number interface{ ~int | ~float64 }

func Half[T Number](v T) T {
	return v * 0.5
}

func TestHalf(t *testing.T) {
	if Half(10) != 5 || Half(3.0) != 1.5 {
		t.Errorf("Half works incorrectly")
	}
}
