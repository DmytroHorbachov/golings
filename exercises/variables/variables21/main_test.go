// variables21
// Make the tests pass!

// I AM NOT DONE
//
// Функция reversed должна вернуть элементы среза в обратном порядке.
// Сейчас она паникует.
// Тренирует: беззнаковый счётчик всегда >= 0, а 0-1 «оборачивается».
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func reversed(items []string) []string {
	out := make([]string, 0, len(items))
	for i := uint(len(items) - 1); i >= 0; i-- {
		out = append(out, items[i])
	}
	return out
}

func TestReversed(t *testing.T) {
	if got := reversed([]string{"a", "b", "c"}); !reflect.DeepEqual(got, []string{"c", "b", "a"}) {
		t.Errorf("reversed(a,b,c) = %v", got)
	}
	if got := reversed(nil); len(got) != 0 {
		t.Errorf("reversed(nil) = %v, want empty", got)
	}
}
