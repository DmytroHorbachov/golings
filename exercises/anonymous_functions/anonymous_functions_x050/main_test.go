// anonymous_functions_x050: Подсчёт по признаку
// Make the tests pass!
// I AM NOT DONE
//
// countBy считает элементы по категориям, заданным литералом.
// Тренирует: функцию-классификатор.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func countBy(nums []int, class func(int) string) map[string]int {
	out := map[string]int{}
	for _, n := range nums {
		out[class(n)] = n
	}
	return out
}

func sizes(nums []int) map[string]int {
	return countBy(nums, func(n int) string {
		if n > 100 {
			return "big"
		}
		return "small"
	})
}

func TestSizes(t *testing.T) {
	if got := sizes([]int{5, 100, 500, 7}); !reflect.DeepEqual(got, map[string]int{"small": 2, "big": 2}) {
		t.Errorf("sizes = %v", got)
	}
}
