// concurrent36
// Make the tests pass!

// I AM NOT DONE
//
// fetch запускает дочернюю операцию с таймаутом, но строит её контекст
// от context.Background, и отмена родителя до неё не доходит.
// Тренирует: производные контексты должны наследовать родительский.
// Сложность: hard
package main_test

import (
	"context"
	"testing"
	"time"
)

func fetch(ctx context.Context) error {
	child, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	<-child.Done()
	return child.Err()
}

func TestFetchCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	res := make(chan error, 1)
	go func() { res <- fetch(ctx) }()
	cancel()
	select {
	case err := <-res:
		if err != context.Canceled {
			t.Errorf("err = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("cancellation did not reach the child context")
	}
}
