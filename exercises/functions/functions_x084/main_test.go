// functions_x084: Сортировка не того среза
// Make the tests pass!
// I AM NOT DONE
//
// topScores должна вернуть отсортированную по убыванию копию, не меняя исходный срез.
// Результат получается перемешанным.
// Тренирует: функция сравнения должна смотреть на тот же срез, который сортируется.
// Сложность: hard
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func topScores(scores []int) []int {
	sorted := make([]int, len(scores))
	copy(sorted, scores)
	sort.Slice(sorted, func(i, j int) bool { return scores[i] > scores[j] })
	return sorted
}

func TestTopScores(t *testing.T) {
	in := []int{3, 9, 1, 7, 5}
	if got := topScores(in); !reflect.DeepEqual(got, []int{9, 7, 5, 3, 1}) {
		t.Errorf("topScores = %v, want [9 7 5 3 1]", got)
	}
	if !reflect.DeepEqual(in, []int{3, 9, 1, 7, 5}) {
		t.Errorf("input modified: %v", in)
	}
}
