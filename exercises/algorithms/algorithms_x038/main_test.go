// algorithms_x038: First Bad Version (первая плохая версия)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: двоичный поиск по предикату. Версии 1..n; начиная с некоторой,
// все версии плохие. Найдите первую плохую, вызывая isBad как можно реже.
// Если плохих нет — верните n+1.
// Сложность: easy. Ожидаемая асимптотика: O(log n) вызовов, O(1) по памяти
package main_test

import "testing"

func firstBad(n int, isBad func(int) bool) int {
	return 0
}

func TestFirstBad(t *testing.T) {
	cases := []struct{ n, bad int }{{5, 4}, {1, 1}, {10, 11}, {2147483647, 2147483646}, {100, 1}}
	for _, c := range cases {
		calls := 0
		got := firstBad(c.n, func(v int) bool { calls++; return v >= c.bad })
		if got != c.bad {
			t.Errorf("firstBad(n=%d) = %d, want %d", c.n, got, c.bad)
		}
		if calls > 40 {
			t.Errorf("firstBad(n=%d) made %d calls", c.n, calls)
		}
	}
}
