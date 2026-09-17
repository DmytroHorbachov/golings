// algorithms_x093: Word Ladder (лестница слов)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: обход в ширину по неявному графу. Найдите длину кратчайшей цепочки
// превращений beginWord → endWord, где на каждом шаге меняется одна буква,
// и каждое промежуточное слово есть в словаре. Если цепочки нет — 0.
// Длина считается в словах, включая начальное и конечное.
// Сложность: hard. Ожидаемая асимптотика: O(N·L²) по времени, O(N·L) по памяти
package main_test

import "testing"

func ladderLength(begin, end string, words []string) int {
	return 0
}

func TestLadderLength(t *testing.T) {
	words := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	if got := ladderLength("hit", "cog", words); got != 5 {
		t.Errorf("ladderLength = %d, want 5", got)
	}
	if got := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}); got != 0 {
		t.Errorf("missing end word should give 0, got %d", got)
	}
	if got := ladderLength("a", "c", []string{"a", "b", "c"}); got != 2 {
		t.Errorf("single letter ladder = %d, want 2", got)
	}
}
