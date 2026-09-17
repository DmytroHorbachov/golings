// arrays_x043: Двоичный поиск в массиве
// Make the tests pass!
// I AM NOT DONE
//
// find ищет значение в отсортированном массиве и возвращает индекс или -1.
// Тренирует: границы lo/hi в двоичном поиске.
// Сложность: medium
package main_test

import "testing"

func find(a [8]int, x int) int {
	lo, hi := 0, len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		switch {
		case a[mid] == x:
			return mid
		case a[mid] < x:
			hi = mid
		default:
			lo = mid
		}
	}
	return -1
}

func TestFind(t *testing.T) {
	a := [8]int{1, 3, 5, 7, 9, 11, 13, 15}
	for i, v := range a {
		if got := find(a, v); got != i {
			t.Errorf("find(%d) = %d, want %d", v, got, i)
		}
	}
	if find(a, 4) != -1 || find(a, 16) != -1 {
		t.Errorf("find should return -1 for missing values")
	}
}
