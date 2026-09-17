// algorithms_x039: Search in Rotated Sorted Array (поиск в сдвинутом массиве)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: модифицированный двоичный поиск. Отсортированный срез различных чисел
// циклически сдвинут. Верните индекс target или -1.
// Сложность: medium. Ожидаемая асимптотика: O(log n) по времени, O(1) по памяти
package main_test

import "testing"

func searchRotated(nums []int, target int) int {
	return 0
}

func TestSearchRotated(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	cases := map[int]int{0: 4, 3: -1, 4: 0, 2: 6, 7: 3}
	for target, want := range cases {
		if got := searchRotated(nums, target); got != want {
			t.Errorf("searchRotated(%d) = %d, want %d", target, got, want)
		}
	}
	if searchRotated([]int{1}, 0) != -1 || searchRotated(nil, 1) != -1 || searchRotated([]int{3, 1}, 1) != 1 {
		t.Errorf("edge cases failed")
	}
}
