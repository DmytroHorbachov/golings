// range_x002: Только индекс
// Make the tests pass!
// I AM NOT DONE
//
// evenPositions возвращает индексы 0, 2, 4, ... для среза.
// Тренирует: range с одной переменной (индексом).
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func evenPositions(s []string) []int {
	var out []int
	for i := range s {
		if i%2 == 1 {
			out = append(out, i)
		}
	}
	return out
}

func TestEvenPositions(t *testing.T) {
	if got := evenPositions([]string{"a", "b", "c", "d", "e"}); !reflect.DeepEqual(got, []int{0, 2, 4}) {
		t.Errorf("evenPositions = %v", got)
	}
}
