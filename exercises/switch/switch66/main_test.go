// switch66
// Make the tests pass!

// I AM NOT DONE
//
// bucket относит число к корзине: "neg" для < 0, "small" для 0..9, "big" для 10+.
// Тренирует: switch без тега как замену цепочке if-else.
// Сложность: easy
package main_test

import "testing"

func bucket(n int) string {
	switch {
	case n < 0:
		return "neg"
	case n <= 10:
		return "small"
	default:
		return "big"
	}
}

func TestBucket(t *testing.T) {
	cases := map[int]string{-1: "neg", 0: "small", 9: "small", 10: "big", 99: "big"}
	for in, want := range cases {
		if got := bucket(in); got != want {
			t.Errorf("bucket(%d) = %s, want %s", in, got, want)
		}
	}
}
