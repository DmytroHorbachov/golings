// maps17
// Make the tests pass!

// I AM NOT DONE
//
// topWord возвращает самое частое слово; при равенстве — первое по алфавиту.
// Тренирует: подсчёт в map и детерминированный выбор при равенстве.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func topWord(text string) string {
	counts := map[string]int{}
	for _, w := range strings.Fields(text) {
		counts[w]++
	}
	best, bestN := "", 0
	for w, n := range counts {
		if n >= bestN {
			best = w
		}
	}
	return best
}

func TestTopWord(t *testing.T) {
	for i := 0; i < 20; i++ {
		if got := topWord("b a c b a d"); got != "a" {
			t.Fatalf("topWord = %q, want a", got)
		}
	}
	if got := topWord("x y y"); got != "y" {
		t.Errorf("topWord = %q, want y", got)
	}
}
