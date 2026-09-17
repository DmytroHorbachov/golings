// range_x009: Первый пробел
// Make the tests pass!
// I AM NOT DONE
//
// firstSpace возвращает байтовую позицию первого пробела или -1.
// Тренирует: индекс в range по строке — позиция байта.
// Сложность: easy
package main_test

import "testing"

func firstSpace(s string) int {
	for i, r := range s {
		if r == ' ' {
			continue
		}
	}
	return -1
}

func TestFirstSpace(t *testing.T) {
	cases := map[string]int{"hello world": 5, "go": -1, " x": 0}
	for in, want := range cases {
		if got := firstSpace(in); got != want {
			t.Errorf("firstSpace(%q) = %d, want %d", in, got, want)
		}
	}
}
