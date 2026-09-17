// switch_x027: Знак через switch
// Make the tests pass!
// I AM NOT DONE
//
// sign возвращает -1, 0 или 1.
// Тренирует: switch без тега с тремя ветками.
// Сложность: easy
package main_test

import "testing"

func sign(x float64) int {
	switch {
	case x > 0:
		return 1
	case x <= 0:
		return -1
	}
	return 0
}

func TestSign(t *testing.T) {
	cases := map[float64]int{2.5: 1, 0: 0, -0.1: -1}
	for in, want := range cases {
		if got := sign(in); got != want {
			t.Errorf("sign(%v) = %d, want %d", in, got, want)
		}
	}
}
