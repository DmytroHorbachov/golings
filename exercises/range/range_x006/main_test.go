// range_x006: Range по каналу
// Make the tests pass!
// I AM NOT DONE
//
// drain собирает все значения из закрытого канала.
// Тренирует: range по каналу продолжается до его закрытия.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func drain(ch <-chan int) []int {
	var out []int
	for v := range ch {
		out = append(out, v-len(out))
	}
	return out
}

func TestDrain(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 5
	ch <- 7
	ch <- 9
	close(ch)
	if got := drain(ch); !reflect.DeepEqual(got, []int{5, 7, 9}) {
		t.Errorf("drain = %v", got)
	}
}
