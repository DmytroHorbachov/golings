// concurrent15
// Make the tests pass!

// I AM NOT DONE
//
// Воркер вызывает wg.Done и в defer, и явно при ошибке.
// Счётчик становится отрицательным, и программа паникует.
// Тренирует: одна горутина — один Done.
// Сложность: hard
package main_test

import (
	"errors"
	"sync"
	"testing"
)

func runChecks(checks []func() error) (fails int) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, c := range checks {
		wg.Add(1)
		go func(c func() error) {
			defer wg.Done()
			if err := c(); err != nil {
				mu.Lock()
				fails++
				mu.Unlock()
				wg.Done()
				return
			}
		}(c)
	}
	wg.Wait()
	return fails
}

func TestRunChecks(t *testing.T) {
	ok := func() error { return nil }
	bad := func() error { return errors.New("x") }
	res := make(chan int, 1)
	go func() {
		defer func() {
			if recover() != nil {
				res <- -1
			}
		}()
		res <- runChecks([]func() error{ok, bad, ok})
	}()
	if got := <-res; got != 1 {
		t.Errorf("fails = %d", got)
	}
}
