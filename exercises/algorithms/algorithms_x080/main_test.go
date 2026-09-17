// algorithms_x080: Subsets (все подмножества)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: перебор с возвратом. Верните все подмножества набора различных чисел.
// Порядок: по возрастанию размера «маски» включений, элементы внутри — как во входе.
// Сложность: medium. Ожидаемая асимптотика: O(n·2^n) по времени, O(n) доп. памяти
package main_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func subsets(nums []int) [][]int {
	return nil
}

func TestSubsets(t *testing.T) {
	_, _ = fmt.Sprint, sort.Strings
	got := subsets([]int{1, 2, 3})
	if len(got) != 8 {
		t.Fatalf("subsets count = %d, want 8", len(got))
	}
	var strs []string
	for _, s := range got {
		strs = append(strs, fmt.Sprint(s))
	}
	sort.Strings(strs)
	want := []string{"[1 2 3]", "[1 2]", "[1 3]", "[1]", "[2 3]", "[2]", "[3]", "[]"}
	if !reflect.DeepEqual(strs, want) {
		t.Errorf("subsets = %v", strs)
	}
	if got := subsets(nil); len(got) != 1 || len(got[0]) != 0 {
		t.Errorf("subsets(nil) = %v", got)
	}
}
