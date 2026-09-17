// slices_x031: Поиск в отсортированном
// Make the tests pass!
// I AM NOT DONE
//
// present проверяет наличие числа в отсортированном срезе через sort.SearchInts.
// Тренирует: sort.SearchInts возвращает позицию вставки.
// Сложность: easy
package main_test

import (
	"sort"
	"testing"
)

func present(s []int, x int) bool {
	i := sort.SearchInts(s, x)
	return i < len(s)
}

func TestPresent(t *testing.T) {
	s := []int{1, 3, 5}
	if !present(s, 3) || present(s, 4) || present(s, 9) {
		t.Errorf("present works incorrectly")
	}
}
