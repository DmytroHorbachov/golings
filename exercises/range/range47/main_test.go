// range47
// Make the tests pass!

// I AM NOT DONE
//
// enumerate нумерует значения из канала. Код не компилируется:
// range по каналу даёт только одно значение.
// Тренирует: в отличие от срезов, у range по каналу нет индекса.
// Сложность: hard
package main_test

import (
	"reflect"
	"strconv"
	"testing"
)

func enumerate(ch <-chan string) []string {
	var out []string
	for i, v := range ch {
		out = append(out, strconv.Itoa(i)+":"+v)
	}
	return out
}

func TestEnumerate(t *testing.T) {
	ch := make(chan string, 2)
	ch <- "a"
	ch <- "b"
	close(ch)
	if got := enumerate(ch); !reflect.DeepEqual(got, []string{"0:a", "1:b"}) {
		t.Errorf("enumerate = %v", got)
	}
}
