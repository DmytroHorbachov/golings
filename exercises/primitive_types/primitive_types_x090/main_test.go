// primitive_types_x090: Отрицательный сдвиг
// Make the tests pass!
// I AM NOT DONE
//
// shiftBy сдвигает число влево на n бит, а при отрицательном n — вправо.
// Сейчас отрицательный n вызывает панику.
// Тренирует: сдвиг на отрицательную величину — паника времени выполнения.
// Сложность: hard
package main_test

import "testing"

func shiftBy(x int, n int) int {
	return x << n
}

func TestShiftBy(t *testing.T) {
	cases := [][3]int{{1, 3, 8}, {16, -2, 4}, {5, 0, 5}}
	for _, c := range cases {
		if got := shiftBy(c[0], c[1]); got != c[2] {
			t.Errorf("shiftBy(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
