// algorithms122
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: двоичный поиск. Верните индекс target в отсортированном срезе
// различных чисел или -1.
// Сложность: easy. Ожидаемая асимптотика: O(log n) по времени, O(1) по памяти
package main_test

import "testing"

func search(nums []int, target int) int {
	return 0
}

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	cases := map[int]int{9: 4, 2: -1, -1: 0, 12: 5, 13: -1}
	for target, want := range cases {
		if got := search(nums, target); got != want {
			t.Errorf("search(%d) = %d, want %d", target, got, want)
		}
	}
	if search(nil, 1) != -1 {
		t.Errorf("search(nil) should be -1")
	}
	big := make([]int, 1000000)
	for i := range big {
		big[i] = i * 2
	}
	if search(big, 1999998) != 999999 || search(big, 7) != -1 {
		t.Errorf("search in big slice failed")
	}
}
