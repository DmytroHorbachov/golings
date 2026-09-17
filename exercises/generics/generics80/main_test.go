// generics80
// Make the tests pass!

// I AM NOT DONE
//
// Percent вычисляет долю part от whole в процентах с дробной частью.
// Для целых типов дробная часть теряется.
// Тренирует: операции над T выполняются в типе T.
// Сложность: hard
package main_test

import "testing"

type Number interface{ ~int | ~float64 }

func Percent[T Number](part, whole T) float64 {
	return float64(part / whole * 100)
}

func TestPercent(t *testing.T) {
	if got := Percent(1, 8); got != 12.5 {
		t.Errorf("Percent(1, 8) = %v, want 12.5", got)
	}
	if got := Percent(1.0, 4.0); got != 25 {
		t.Errorf("Percent(1.0, 4.0) = %v", got)
	}
}
