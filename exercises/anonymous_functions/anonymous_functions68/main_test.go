// anonymous_functions68
// Make the tests pass!

// I AM NOT DONE
//
// Литерал должен добавлять бонус к общему счёту, но его параметр назван
// так же, как внешняя переменная, и изменения теряются.
// Тренирует: параметр литерала затеняет захваченную переменную.
// Сложность: hard
package main_test

import "testing"

func play(bonuses []int) int {
	score := 0
	addBonus := func(score int) {
		score += score
	}
	for _, b := range bonuses {
		addBonus(b)
	}
	return score
}

func TestPlay(t *testing.T) {
	if got := play([]int{5, 10}); got != 15 {
		t.Errorf("play = %d, want 15", got)
	}
}
