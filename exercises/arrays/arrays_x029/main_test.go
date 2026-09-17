// arrays_x029: Счётчик голосов
// Make the tests pass!
// I AM NOT DONE
//
// vote увеличивает число голосов за кандидата с номером i.
// Тренирует: изменение элемента массива через указатель.
// Сложность: easy
package main_test

import "testing"

func vote(votes *[3]int, i int) {
	votes[i] = i
}

func TestVote(t *testing.T) {
	var v [3]int
	for _, c := range []int{0, 2, 2, 1, 2} {
		vote(&v, c)
	}
	if v != [3]int{1, 1, 3} {
		t.Errorf("votes = %v, want [1 1 3]", v)
	}
}
