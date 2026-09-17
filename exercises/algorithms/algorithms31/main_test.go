// algorithms31
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: двумерное ДП. Найдите минимальное число операций (вставка,
// удаление, замена символа), превращающих одну строку в другую.
// Сложность: hard. Ожидаемая асимптотика: O(n·m) по времени, O(min(n, m)) по памяти
package main_test

import "testing"

func minDistance(a, b string) int {
	return 0
}

func TestMinDistance(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "abc", 3},
		{"abc", "", 3},
		{"", "", 0},
		{"same", "same", 0},
	}
	for _, c := range cases {
		if got := minDistance(c.a, c.b); got != c.want {
			t.Errorf("minDistance(%q, %q) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
