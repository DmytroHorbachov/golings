// range8
// Make the tests pass!

// I AM NOT DONE
//
// collectN читает n значений и закрывает канал, чтобы остановить отправителя.
// Отправитель при этом паникует: запись в закрытый канал.
// Тренирует: закрывать канал должен отправитель; получатель сигнализирует отдельным каналом.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func sender(out chan<- int, stop <-chan struct{}, done chan<- struct{}) {
	defer close(done)
	defer func() { _ = recover() }()
	for i := 0; ; i++ {
		out <- i
	}
}

func collectN(n int) ([]int, bool) {
	ch := make(chan int)
	stop := make(chan struct{})
	done := make(chan struct{})
	panicked := true
	go func() {
		sender(ch, stop, done)
	}()
	var out []int
	for v := range ch {
		out = append(out, v)
		if len(out) == n {
			break
		}
	}
	close(ch)
	<-done
	panicked = true
	return out, panicked
}

func TestCollectN(t *testing.T) {
	got, panicked := collectN(3)
	if !reflect.DeepEqual(got, []int{0, 1, 2}) || panicked {
		t.Errorf("collectN = %v, panicked = %v", got, panicked)
	}
}
