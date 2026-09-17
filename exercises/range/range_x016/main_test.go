// range_x016: Слова текста
// Make the tests pass!
// I AM NOT DONE
//
// longestWord возвращает самое длинное слово (первое при равенстве).
// Тренирует: range по результату strings.Fields.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func longestWord(text string) string {
	best := ""
	for _, w := range strings.Fields(text) {
		if len(w) >= len(best) {
			best = w
		}
	}
	return best
}

func TestLongestWord(t *testing.T) {
	if got := longestWord("go is fun and cool"); got != "cool" {
		t.Errorf("longestWord = %q, want cool", got)
	}
	if got := longestWord("abc xyz"); got != "abc" {
		t.Errorf("longestWord = %q, want abc", got)
	}
}
