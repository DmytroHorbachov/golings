// algorithms36
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: быстрое возведение в степень. Вычислите x в степени n
// (n может быть отрицательным) без math.Pow.
// Сложность: medium. Ожидаемая асимптотика: O(log n) по времени, O(1) по памяти
package main_test

import (
	"math"
	"testing"
)

func myPow(x float64, n int) float64 {
	return 0
}

func TestMyPow(t *testing.T) {
	cases := []struct {
		x    float64
		n    int
		want float64
	}{
		{2, 10, 1024},
		{2, -2, 0.25},
		{5, 0, 1},
		{1, 1000000, 1},
		{-2, 3, -8},
	}
	for _, c := range cases {
		if got := myPow(c.x, c.n); math.Abs(got-c.want) > 1e-9 {
			t.Errorf("myPow(%v, %d) = %v, want %v", c.x, c.n, got, c.want)
		}
	}
}
