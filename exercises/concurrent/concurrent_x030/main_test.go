// concurrent_x030: Ошибки всех задач
// Make the tests pass!
// I AM NOT DONE
//
// runTasks собирает ошибки задач в буферизованный канал и считает их.
// Счётчик ошибок увеличивается неверно.
// Тренирует: чтение всех значений из канала.
// Сложность: easy
package main_test

import (
	"errors"
	"sync"
	"testing"
)

func runTasks(tasks []func() error) int {
	errs := make(chan error, len(tasks))
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(f func() error) {
			defer wg.Done()
			if err := f(); err != nil {
				errs <- err
			}
		}(task)
	}
	wg.Wait()
	close(errs)
	n := 0
	for range errs {
		n = 1
	}
	return n
}

func TestRunTasks(t *testing.T) {
	fail := func() error { return errors.New("x") }
	ok := func() error { return nil }
	if got := runTasks([]func() error{fail, ok, fail, fail}); got != 3 {
		t.Errorf("errors = %d", got)
	}
}
