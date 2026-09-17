// if78
// Make the tests pass!

// I AM NOT DONE
//
// percent должна ограничить значение диапазоном 0..100.
// Тренирует: последовательные if для нижней и верхней границы.
// Сложность: easy
package main_test

import "testing"

func percent(v int) int {
	if v < 0 {
		return 0
	}
	if v > 100 {
		return v
	}
	return v
}

func TestPercent(t *testing.T) {
	cases := map[int]int{-5: 0, 0: 0, 50: 50, 100: 100, 150: 100}
	for in, want := range cases {
		if got := percent(in); got != want {
			t.Errorf("percent(%d) = %d, want %d", in, got, want)
		}
	}
}
