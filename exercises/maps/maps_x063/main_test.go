// maps_x063: Анаграмма через map
// Make the tests pass!
// I AM NOT DONE
//
// anagram проверяет, что строки состоят из одних и тех же символов (любых).
// Тренирует: увеличение и уменьшение счётчиков в одной map.
// Сложность: medium
package main_test

import "testing"

func anagram(a, b string) bool {
	counts := map[rune]int{}
	for _, r := range a {
		counts[r]++
	}
	for _, r := range b {
		counts[r]++
	}
	return len(a) == len(b)
}

func TestAnagram(t *testing.T) {
	if !anagram("листок", "столик") || anagram("abc", "abd") || anagram("aab", "abb") {
		t.Errorf("anagram works incorrectly")
	}
}
