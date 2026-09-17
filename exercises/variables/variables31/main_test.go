// variables31
// Make the tests pass!

// I AM NOT DONE
//
// Функция round должна округлять до ближайшего целого, половинки — от нуля.
// Для отрицательных чисел текущая формула ошибается.
// Тренирует: преобразование float в int отбрасывает дробную часть к нулю.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func round(x float64) int {
	_ = math.Round
	return int(x + 0.5)
}

func TestRound(t *testing.T) {
	cases := map[float64]int{2.4: 2, 2.5: 3, -2.4: -2, -2.5: -3, -0.6: -1, 0: 0}
	for in, want := range cases {
		if got := round(in); got != want {
			t.Errorf("round(%v) = %d, want %d", in, got, want)
		}
	}
}
