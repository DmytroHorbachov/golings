// arrays_x084: sort.Slice и массив
// Make the tests pass!
// I AM NOT DONE
//
// sortScores сортирует массив очков по убыванию. Сейчас код паникует:
// sort.Slice получает массив, а не срез.
// Тренирует: sort.Slice принимает interface{} и работает только со срезами.
// Сложность: hard
package main_test

import (
	"sort"
	"testing"
)

func sortScores(a *[5]int) {
	sort.Slice(*a, func(i, j int) bool { return a[i] > a[j] })
}

func TestSortScores(t *testing.T) {
	a := [5]int{3, 9, 1, 7, 5}
	sortScores(&a)
	if a != [5]int{9, 7, 5, 3, 1} {
		t.Errorf("sortScores = %v", a)
	}
}
