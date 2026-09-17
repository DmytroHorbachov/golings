// maps73
// Make the tests pass!

// I AM NOT DONE
//
// runeFreq считает символы строки, включая кириллицу.
// Тренирует: map[rune]int и range по строке.
// Сложность: easy
package main_test

import "testing"

func runeFreq(s string) map[rune]int {
	m := map[rune]int{}
	for _, r := range s {
		m[r] = 1
	}
	return m
}

func TestRuneFreq(t *testing.T) {
	m := runeFreq("αβαβ")
	if m['α'] != 2 || m['β'] != 2 || len(m) != 2 {
		t.Errorf("runeFreq = %v", m)
	}
}
