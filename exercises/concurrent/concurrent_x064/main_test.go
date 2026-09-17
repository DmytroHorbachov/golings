// concurrent_x064: Дедупликация потока
// Make the tests pass!
// I AM NOT DONE
//
// dedupe читает поток и пропускает дальше только первые вхождения значений.
// Тренирует: горутина-этап со своим состоянием.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
	"time"
)

func dedupe(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		seen := map[string]bool{}
		for v := range in {
			out <- v
		}
	}()
	return out
}

func TestDedupe(t *testing.T) {
	in := make(chan string, 5)
	for _, s := range []string{"a", "b", "a", "c", "b"} {
		in <- s
	}
	close(in)
	res := make(chan []string, 1)
	go func() {
		var got []string
		for v := range dedupe(in) {
			got = append(got, v)
		}
		res <- got
	}()
	select {
	case got := <-res:
		if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
			t.Errorf("dedupe = %v", got)
		}
	case <-time.After(time.Second):
		t.Fatal("dedupe is stuck")
	}
}
