// anonymous_functions_x073: Фабрика в цикле
// Make the tests pass!
// I AM NOT DONE
//
// processAll должна выдать всем записям последовательные номера, но все
// получают 1: генератор создаётся заново для каждой записи.
// Тренирует: состояние замыкания живёт, пока жив сам литерал.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func counter() func() int {
	n := 0
	return func() int { n++; return n }
}

func processAll(names []string) []int {
	var ids []int
	for range names {
		next := counter()
		ids = append(ids, next())
	}
	return ids
}

func TestProcessAll(t *testing.T) {
	if got := processAll([]string{"a", "b", "c"}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("processAll = %v", got)
	}
}
