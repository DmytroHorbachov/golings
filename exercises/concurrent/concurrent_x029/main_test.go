// concurrent_x029: Канал таймера
// Make the tests pass!
// I AM NOT DONE
//
// waitTimer ждёт срабатывания таймера через его канал.
// Тренирует: time.NewTimer и поле C.
// Сложность: easy
package main_test

import (
	"testing"
	"time"
)

func waitTimer(d time.Duration) bool {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-make(chan time.Time):
		return true
	case <-time.After(time.Second):
		return false
	}
}

func TestWaitTimer(t *testing.T) {
	if !waitTimer(10 * time.Millisecond) {
		t.Errorf("timer did not fire")
	}
}
