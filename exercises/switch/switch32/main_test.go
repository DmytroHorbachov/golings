// switch32
// Make the tests pass!

// I AM NOT DONE
//
// ratingStars переводит оценку 0–100 в звёзды: 90+ — 3, 70+ — 2, 50+ — 1, иначе 0.
// Тренирует: порядок веток в switch без тега.
// Сложность: easy
package main_test

import "testing"

func ratingStars(score int) int {
	switch {
	case score >= 90:
		return 3
	case score > 70:
		return 2
	case score >= 50:
		return 1
	}
	return 0
}

func TestRatingStars(t *testing.T) {
	cases := map[int]int{95: 3, 90: 3, 70: 2, 69: 1, 50: 1, 10: 0}
	for in, want := range cases {
		if got := ratingStars(in); got != want {
			t.Errorf("ratingStars(%d) = %d, want %d", in, got, want)
		}
	}
}
