// slices_x034: Минимум среза
// Make the tests pass!
// I AM NOT DONE
//
// minOf возвращает минимум непустого среза.
// Тренирует: сравнение элементов при обходе.
// Сложность: easy
package main_test

import "testing"

func minOf(s []float64) float64 {
	m := s[0]
	for _, v := range s[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func TestMinOf(t *testing.T) {
	if got := minOf([]float64{2.5, -1, 7}); got != -1 {
		t.Errorf("minOf = %v", got)
	}
}
