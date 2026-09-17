// arrays_x004: Нулевой массив
// Make the tests pass!
// I AM NOT DONE
//
// emptyScores должна вернуть массив из трёх нулей.
// Тренирует: нулевое значение массива.
// Сложность: easy
package main_test

import "testing"

func emptyScores() [3]int {
	var scores [3]int
	scores[1] = 1
	return scores
}

func TestEmptyScores(t *testing.T) {
	if got := emptyScores(); got != [3]int{} {
		t.Errorf("emptyScores = %v, want [0 0 0]", got)
	}
}
