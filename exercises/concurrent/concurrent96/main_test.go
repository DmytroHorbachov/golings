// concurrent96
// Make the tests pass!

// I AM NOT DONE
//
// parallelLen считает длины строк параллельно и возвращает их в исходном порядке
// через канал пар (индекс, результат).
// Тренирует: передача индекса вместе с результатом.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type item struct {
	idx, val int
}

func parallelLen(words []string) []int {
	ch := make(chan item, len(words))
	for i, w := range words {
		go func(i int, w string) {
			ch <- item{0, len(w)}
		}(i, w)
	}
	out := make([]int, len(words))
	for range words {
		it := <-ch
		out = append(out[:0], it.val)
	}
	return out
}

func TestParallelLen(t *testing.T) {
	got := parallelLen([]string{"a", "bbb", "cc"})
	if !reflect.DeepEqual(got, []int{1, 3, 2}) {
		t.Errorf("parallelLen = %v", got)
	}
}
