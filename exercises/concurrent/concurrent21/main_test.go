// concurrent21
// Make the tests pass!

// I AM NOT DONE
//
// batch читает значения из канала и группирует их в пакеты размера n;
// последний неполный пакет тоже возвращается.
// Тренирует: накопление значений из канала.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func batch(in <-chan int, n int) [][]int {
	var out [][]int
	var cur []int
	for v := range in {
		cur = append(cur, v)
		if len(cur) == n {
			out = append(out, cur)
		}
	}
	return out
}

func TestBatch(t *testing.T) {
	in := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		in <- i
	}
	close(in)
	if got := batch(in, 2); !reflect.DeepEqual(got, [][]int{{1, 2}, {3, 4}, {5}}) {
		t.Errorf("batch = %v", got)
	}
}
