// algorithms_x082: Combination Sum (комбинации с повторами)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: перебор с возвратом. Из различных положительных чисел составьте все
// комбинации с суммой target; каждое число можно брать сколько угодно раз.
// Комбинации не должны повторяться (числа внутри — по неубыванию).
// Сложность: medium. Ожидаемая асимптотика: O(n^(t/min)) по времени, O(t/min) по памяти
package main_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	return nil
}

func TestCombinationSum(t *testing.T) {
	_, _ = fmt.Sprint, sort.Strings
	got := combinationSum([]int{2, 3, 6, 7}, 7)
	var strs []string
	for _, c := range got {
		strs = append(strs, fmt.Sprint(c))
	}
	sort.Strings(strs)
	if !reflect.DeepEqual(strs, []string{"[2 2 3]", "[7]"}) {
		t.Errorf("combinationSum = %v", strs)
	}
	if got := combinationSum([]int{2}, 1); len(got) != 0 {
		t.Errorf("no combinations expected, got %v", got)
	}
	if got := combinationSum([]int{1}, 0); len(got) != 1 || len(got[0]) != 0 {
		t.Errorf("zero target = %v", got)
	}
}
