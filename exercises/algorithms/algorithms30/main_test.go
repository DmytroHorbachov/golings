// algorithms30
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: сортировка и два указателя. Найдите все уникальные тройки чисел
// с суммой 0. Каждая тройка упорядочена по возрастанию, тройки — лексикографически.
// Сложность: medium. Ожидаемая асимптотика: O(n²) по времени, O(1) доп. памяти (без учёта сортировки и ответа)
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func threeSum(nums []int) [][3]int {
	return nil
}

func TestThreeSum(t *testing.T) {
	_ = sort.Ints
	got := threeSum([]int{-1, 0, 1, 2, -1, -4})
	if !reflect.DeepEqual(got, [][3]int{{-1, -1, 2}, {-1, 0, 1}}) {
		t.Errorf("threeSum = %v", got)
	}
	if got := threeSum([]int{0, 0, 0, 0}); !reflect.DeepEqual(got, [][3]int{{0, 0, 0}}) {
		t.Errorf("threeSum(zeros) = %v", got)
	}
	if got := threeSum([]int{0, 1, 1}); len(got) != 0 {
		t.Errorf("threeSum(no triples) = %v", got)
	}
	if got := threeSum(nil); len(got) != 0 {
		t.Errorf("threeSum(nil) = %v", got)
	}
}
