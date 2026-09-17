// slices72
// Make the tests pass!

// I AM NOT DONE
//
// insert вставляет x в позицию i. Результат содержит дубликаты вместо сдвига.
// Тренирует: append(append(s[:i], x), s[i:]...) перезаписывает s[i] до копирования хвоста.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func insert(s []int, i, x int) []int {
	return append(append(s[:i], x), s[i:]...)
}

func TestInsert(t *testing.T) {
	s := make([]int, 3, 10)
	copy(s, []int{1, 2, 4})
	if got := insert(s, 2, 3); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("insert = %v, want [1 2 3 4]", got)
	}
}
