// range_x088: Удаление обменом с последним
// Make the tests pass!
// I AM NOT DONE
//
// removeExpired удаляет истёкшие токены обменом с последним элементом.
// Во время range элемент, переставленный на место i, не проверяется.
// Тренирует: range не учитывает перестановки за текущей позицией.
// Сложность: hard
package main_test

import (
	"sort"
	"testing"
)

func removeExpired(tokens []int, now int) []int {
	for i, exp := range tokens {
		if i < len(tokens) && exp < now {
			last := len(tokens) - 1
			tokens[i] = tokens[last]
			tokens = tokens[:last]
		}
	}
	return tokens
}

func TestRemoveExpired(t *testing.T) {
	got := removeExpired([]int{5, 20, 1, 30, 2}, 10)
	sort.Ints(got)
	if len(got) != 2 || got[0] != 20 || got[1] != 30 {
		t.Errorf("removeExpired = %v, want [20 30]", got)
	}
}
