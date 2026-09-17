// concurrent_x042: Первая ошибка отменяет остальных
// Make the tests pass!
// I AM NOT DONE
//
// runAll выполняет задачи параллельно; первая ошибка отменяет контекст
// остальных задач и возвращается.
// Тренирует: context.WithCancel и sync.Once для первой ошибки.
// Сложность: medium
package main_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func runAll(tasks []func(context.Context) error) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	var once sync.Once
	var first error
	for _, task := range tasks {
		wg.Add(1)
		go func(f func(context.Context) error) {
			defer wg.Done()
			if err := f(ctx); err != nil {
				first = err
			}
		}(task)
	}
	wg.Wait()
	return first
}

func TestRunAll(t *testing.T) {
	boom := errors.New("boom")
	failing := func(ctx context.Context) error { return boom }
	waiting := func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(3 * time.Second):
			return nil
		}
	}
	start := time.Now()
	err := runAll([]func(context.Context) error{waiting, failing, waiting})
	if err != boom || time.Since(start) > time.Second {
		t.Errorf("err = %v after %v", err, time.Since(start))
	}
}
