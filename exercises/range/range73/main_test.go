// range73
// Make the tests pass!

// I AM NOT DONE
//
// longestRun возвращает символ с самой длинной серией подряд и длину серии.
// Тренирует: range по строке с текущей и лучшей серией.
// Сложность: medium
package main_test

import "testing"

func longestRun(s string) (rune, int) {
	var best, prev rune
	bestN, cur := 0, 0
	for _, r := range s {
		if r == prev {
			cur++
		}
		if cur > bestN {
			best, bestN = r, cur
		}
	}
	return best, bestN
}

func TestLongestRun(t *testing.T) {
	if r, n := longestRun("aabbbbccж"); r != 'b' || n != 4 {
		t.Errorf("longestRun = %c, %d", r, n)
	}
	if r, n := longestRun("ёёё"); r != 'ё' || n != 3 {
		t.Errorf("longestRun(ёёё) = %c, %d", r, n)
	}
}
