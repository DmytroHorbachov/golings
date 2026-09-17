// concurrent_x094: Таймаут в цикле
// Make the tests pass!
// I AM NOT DONE
//
// listen должна завершиться через total после начала, даже если сообщения
// продолжают приходить. time.After в каждой итерации перезапускает таймер.
// Тренирует: общий таймаут создаётся один раз.
// Сложность: hard
package main_test

import (
	"testing"
	"time"
)

func listen(msgs <-chan int, total time.Duration) int {
	n := 0
	for {
		select {
		case <-msgs:
			n++
		case <-time.After(total):
			return n
		}
	}
}

func TestListen(t *testing.T) {
	msgs := make(chan int)
	stop := make(chan struct{})
	defer close(stop)
	go func() {
		for {
			select {
			case msgs <- 1:
				time.Sleep(5 * time.Millisecond)
			case <-stop:
				return
			}
		}
	}()
	res := make(chan int, 1)
	go func() { res <- listen(msgs, 100*time.Millisecond) }()
	select {
	case n := <-res:
		if n == 0 {
			t.Errorf("no messages received")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("listen never timed out")
	}
}
