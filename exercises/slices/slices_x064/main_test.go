// slices_x064: Все позиции
// Make the tests pass!
// I AM NOT DONE
//
// allIndexes возвращает все индексы, где встречается x.
// Тренирует: накопление индексов в срез.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func allIndexes(s []string, x string) []int {
	var out []int
	for i, v := range s {
		if v == x {
			out = append(out, len(out))
			break
		}
	}
	return out
}

func TestAllIndexes(t *testing.T) {
	if got := allIndexes([]string{"a", "b", "a", "a"}, "a"); !reflect.DeepEqual(got, []int{0, 2, 3}) {
		t.Errorf("allIndexes = %v", got)
	}
	if got := allIndexes([]string{"a"}, "z"); got != nil {
		t.Errorf("allIndexes(missing) = %v", got)
	}
}
