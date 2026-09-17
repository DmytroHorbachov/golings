// maps_x050: Элемент большинства
// Make the tests pass!
// I AM NOT DONE
//
// majority возвращает элемент, встречающийся больше n/2 раз, или false.
// Тренирует: подсчёт и сравнение с порогом.
// Сложность: medium
package main_test

import "testing"

func majority(nums []int) (int, bool) {
	counts := map[int]int{}
	for _, v := range nums {
		counts[v]++
		if counts[v] >= len(nums)/2 {
			return v, true
		}
	}
	return nums[0], true
}

func TestMajority(t *testing.T) {
	if v, ok := majority([]int{2, 2, 1, 1, 1, 2, 2}); !ok || v != 2 {
		t.Errorf("majority = %d, %v", v, ok)
	}
	if _, ok := majority([]int{1, 2, 3, 1}); ok {
		t.Errorf("majority([1 2 3 1]) should not exist")
	}
}
