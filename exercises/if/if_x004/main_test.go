// if_x004: Знак числа
// Make the tests pass!
// I AM NOT DONE
//
// sign должна вернуть -1, 0 или 1 в зависимости от знака числа.
// Тренирует: цепочку if / else if / else.
// Сложность: easy
package main_test

import "testing"

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x == 0 {
		return 0
	} else {
		return 1
	}
}

func TestSign(t *testing.T) {
	cases := map[int]int{5: 1, 0: 0, -3: -1}
	for in, want := range cases {
		if got := sign(in); got != want {
			t.Errorf("sign(%d) = %d, want %d", in, got, want)
		}
	}
}
