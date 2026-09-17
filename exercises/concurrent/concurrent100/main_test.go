// concurrent100
// Make the tests pass!

// I AM NOT DONE
//
// emit отправляет события в канал; когда получатель ушёл и контекст отменён,
// отправка блокируется навсегда.
// Тренирует: блокирующие отправки тоже должны слушать ctx.Done().
// Сложность: hard
package main_test

import (
	"context"
	"testing"
	"time"
)

func emit(ctx context.Context, out chan<- int, n int) error {
	for i := 0; i < n; i++ {
		out <- i
	}
	return nil
}

func TestEmit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	out := make(chan int)
	res := make(chan error, 1)
	go func() { res <- emit(ctx, out, 10) }()
	<-out
	<-out
	cancel()
	select {
	case err := <-res:
		if err != context.Canceled {
			t.Errorf("err = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("emit is blocked after cancellation")
	}
}
