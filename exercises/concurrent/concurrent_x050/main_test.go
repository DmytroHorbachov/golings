// concurrent_x050: Остановка тикера
// Make the tests pass!
// I AM NOT DONE
//
// poll выполняет проверку по тикеру заданное число раз и останавливает тикер.
// Тренирует: time.Ticker и его остановку.
// Сложность: medium
package main_test

import (
	"testing"
	"time"
)

func poll(times int, check func()) {
	ticker := time.NewTicker(2 * time.Millisecond)
	for range ticker.C {
		check()
	}
}

func TestPoll(t *testing.T) {
	n := 0
	done := make(chan struct{})
	go func() {
		poll(3, func() { n++ })
		close(done)
	}()
	select {
	case <-done:
		if n != 3 {
			t.Errorf("checks = %d", n)
		}
	case <-time.After(time.Second):
		t.Fatal("poll never stopped")
	}
}
