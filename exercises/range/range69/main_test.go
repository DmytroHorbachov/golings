// range69
// Make the tests pass!

// I AM NOT DONE
//
// totalScore должна сложить очки игроков по map «номер игрока -> очки».
// Сумма оказывается суммой номеров.
// Тренирует: range по map с одной переменной выдаёт ключи.
// Сложность: hard
package main_test

import "testing"

func totalScore(scores map[int]int) int {
	sum := 0
	for v := range scores {
		sum += v
	}
	return sum
}

func TestTotalScore(t *testing.T) {
	if got := totalScore(map[int]int{1: 100, 2: 250, 3: 50}); got != 400 {
		t.Errorf("totalScore = %d, want 400", got)
	}
}
