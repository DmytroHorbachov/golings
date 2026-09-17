// algorithms_x120: Distinct Subsequences (число вхождений подпоследовательностью)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: двумерное ДП. Посчитайте, сколькими способами можно получить
// строку t как подпоследовательность строки s.
// Сложность: hard. Ожидаемая асимптотика: O(n·m) по времени, O(m) по памяти
package main_test

import "testing"

func numDistinct(s, t string) int {
	return 0
}

func TestNumDistinct(t *testing.T) {
	cases := []struct {
		s, t string
		want int
	}{
		{"rabbbit", "rabbit", 3},
		{"babgbag", "bag", 5},
		{"abc", "", 1},
		{"", "a", 0},
		{"aaa", "aa", 3},
	}
	for _, c := range cases {
		if got := numDistinct(c.s, c.t); got != c.want {
			t.Errorf("numDistinct(%q, %q) = %d, want %d", c.s, c.t, got, c.want)
		}
	}
}
