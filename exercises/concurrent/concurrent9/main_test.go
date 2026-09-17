// concurrent9
// Make the tests pass!

// I AM NOT DONE
//
// readAll читает значения, пока канал не закрыт, используя форму v, ok.
// Тренирует: признак ok при чтении из канала.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func readAll(ch <-chan int) []int {
	var out []int
	for {
		v, ok := <-ch
		if ok {
			return out
		}
		out = append(out, v)
	}
}

func TestReadAll(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 4
	ch <- 5
	close(ch)
	if got := readAll(ch); !reflect.DeepEqual(got, []int{4, 5}) {
		t.Errorf("readAll = %v", got)
	}
}
