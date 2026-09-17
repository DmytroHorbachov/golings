// generics_x076: Преобразование в float64
// Make the tests pass!
// I AM NOT DONE
//
// Ratio делит два значения параметра типа как float64.
// Код не компилируется: T ограничен any, а преобразование к float64
// допустимо только для числовых типов.
// Тренирует: конверсии зависят от набора типов ограничения.
// Сложность: hard
package main_test

import "testing"

type Number interface{ ~int | ~int64 | ~float64 }

func Ratio[T any](a, b T) float64 {
	return float64(a) / float64(b)
}

func TestRatio(t *testing.T) {
	if Ratio(1, 4) != 0.25 || Ratio(3.0, 2.0) != 1.5 {
		t.Errorf("Ratio works incorrectly")
	}
}
