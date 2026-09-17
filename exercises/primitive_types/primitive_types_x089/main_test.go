// primitive_types_x089: Корзина для отрицательных
// Make the tests pass!
// I AM NOT DONE
//
// bucket возвращает номер интервала шириной size, куда попадает x:
// [0, size) — 0, [-size, 0) — -1.
// Тренирует: int(x) усекает к нулю, а не округляет вниз.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func bucket(x, size float64) int {
	return int(x / size)
}

func TestBucket(t *testing.T) {
	_ = math.Floor
	cases := map[float64]int{0: 0, 9.9: 0, 10: 1, -0.5: -1, -10: -1, -10.1: -2}
	for in, want := range cases {
		if got := bucket(in, 10); got != want {
			t.Errorf("bucket(%v, 10) = %d, want %d", in, got, want)
		}
	}
}
