// concurrent_x071: Отправка в закрытый канал
// Make the tests pass!
// I AM NOT DONE
//
// Несколько производителей пишут в один канал; первый закончивший закрывает его,
// и остальные паникуют.
// Тренирует: закрывать канал должен тот, кто знает, что отправителей больше нет.
// Сложность: hard
package main_test

import (
	"sync"
	"testing"
	"time"
)

func produceAll(workers, each int) int {
	out := make(chan int)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			defer close(out)
			for i := 0; i < each; i++ {
				time.Sleep(time.Duration(w) * time.Millisecond)
				out <- 1
			}
		}(w)
	}
	total := 0
	for v := range out {
		total += v
	}
	return total
}

func TestProduceAll(t *testing.T) {
	res := make(chan int, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				res <- -1
			}
		}()
		res <- produceAll(3, 4)
	}()
	select {
	case got := <-res:
		if got != 12 {
			t.Errorf("total = %d, want 12", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("produceAll is stuck")
	}
}
