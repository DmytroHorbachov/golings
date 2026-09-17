// range_x067: Разбиение по разделителю
// Make the tests pass!
// I AM NOT DONE
//
// splitOn делит срез чисел на части по значению-разделителю sep (разделитель не входит в части).
// Тренирует: range с накоплением текущей части.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func splitOn(s []int, sep int) [][]int {
	var out [][]int
	cur := []int{}
	for _, v := range s {
		if v == sep {
			continue
		}
		cur = append(cur, v)
	}
	return out
}

func TestSplitOn(t *testing.T) {
	got := splitOn([]int{1, 2, 0, 3, 0, 0, 4}, 0)
	want := [][]int{{1, 2}, {3}, {}, {4}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitOn = %v, want %v", got, want)
	}
}
