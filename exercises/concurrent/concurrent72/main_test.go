// concurrent72
// Make the tests pass!

// I AM NOT DONE
//
// checkAll проверяет адреса параллельно и сохраняет ошибку в общую переменную
// без синхронизации.
// Тренирует: запись в общую переменную из горутин требует защиты.
// Сложность: hard
package main_test

import (
	"errors"
	"strings"
	"sync"
	"testing"
)

func checkAll(hosts []string) error {
	var mu sync.Mutex
	var firstErr error
	var wg sync.WaitGroup
	for _, h := range hosts {
		wg.Add(1)
		go func(h string) {
			defer wg.Done()
			if strings.HasPrefix(h, "bad") {
				firstErr = errors.New("unreachable " + h)
			}
		}(h)
	}
	wg.Wait()
	_ = &mu
	return firstErr
}

func TestCheckAll(t *testing.T) {
	hosts := []string{"a", "bad1", "bad2", "b", "bad3", "bad4"}
	if err := checkAll(hosts); err == nil {
		t.Errorf("expected an error")
	}
}
