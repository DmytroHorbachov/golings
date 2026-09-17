// primitive_types_x053: Целый квадратный корень
// Make the tests pass!
// I AM NOT DONE
//
// isqrt возвращает наибольшее целое r, такое что r*r <= n.
// Для больших n результат math.Sqrt может быть неточным.
// Тренирует: преобразования float64/int и коррекцию погрешности.
// Сложность: medium
package main_test

import (
	"math"
	"testing"
)

func isqrt(n int64) int64 {
	r := int64(math.Sqrt(float64(n)))
	return r
}

func TestIsqrt(t *testing.T) {
	cases := map[int64]int64{0: 0, 15: 3, 16: 4, 999999999999999999: 999999999, 1000000000000000000: 1000000000}
	for in, want := range cases {
		if got := isqrt(in); got != want {
			t.Errorf("isqrt(%d) = %d, want %d", in, got, want)
		}
	}
}
