// algorithms133
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: одномерное ДП. Сколькими способами можно подняться на n ступеней,
// делая шаги по 1 или 2 ступени? Для n = 0 ответ 1.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func climbStairs(n int) int {
	return 0
}

func TestClimbStairs(t *testing.T) {
	cases := map[int]int{0: 1, 1: 1, 2: 2, 3: 3, 5: 8, 45: 1836311903}
	for n, want := range cases {
		if got := climbStairs(n); got != want {
			t.Errorf("climbStairs(%d) = %d, want %d", n, got, want)
		}
	}
}
