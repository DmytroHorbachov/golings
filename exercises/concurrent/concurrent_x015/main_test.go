// concurrent_x015: Вызов cancel
// Make the tests pass!
// I AM NOT DONE
//
// stopAll должна отменить контекст, чтобы воркер завершился.
// Тренирует: context.WithCancel.
// Сложность: easy
package main_test

import (
	"context"
	"testing"
	"time"
)

func stopAll() bool {
	ctx, cancel := context.WithCancel(context.Background())
	stopped := make(chan struct{})
	go func() {
		<-ctx.Done()
		close(stopped)
	}()
	_ = cancel
	select {
	case <-stopped:
		return true
	case <-time.After(time.Second):
		return false
	}
}

func TestStopAll(t *testing.T) {
	if !stopAll() {
		t.Errorf("worker was not stopped")
	}
}
