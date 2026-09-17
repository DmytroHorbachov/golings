// maps14
// Make the tests pass!

// I AM NOT DONE
//
// countPairs считает, сколько раз встречается каждая упорядоченная пара соседних слов.
// Тренирует: структура из двух полей как ключ.
// Сложность: medium
package main_test

import "testing"

type bigram struct{ a, b string }

func countPairs(words []string) map[bigram]int {
	m := map[bigram]int{}
	for i := range words {
		m[bigram{words[i], words[i]}]++
	}
	return m
}

func TestCountPairs(t *testing.T) {
	m := countPairs([]string{"to", "be", "or", "not", "to", "be"})
	if m[bigram{"to", "be"}] != 2 || m[bigram{"be", "or"}] != 1 || len(m) != 4 {
		t.Errorf("countPairs = %v", m)
	}
}
