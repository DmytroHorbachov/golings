// if_x029: Число в списке
// Make the tests pass!
// I AM NOT DONE
//
// contains должна вернуть true, если x есть в срезе.
// Сейчас функция выходит слишком рано.
// Тренирует: return внутри if в цикле.
// Сложность: easy
package main_test

import "testing"

func contains(nums []int, x int) bool {
	for _, n := range nums {
		if n == x {
			return true
		}
		return false
	}
	return false
}

func TestContains(t *testing.T) {
	nums := []int{4, 8, 15, 16}
	if !contains(nums, 15) || !contains(nums, 4) {
		t.Errorf("contains should find 15 and 4")
	}
	if contains(nums, 42) || contains(nil, 1) {
		t.Errorf("contains should not find missing values")
	}
}
