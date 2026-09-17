// anonymous_functions65
// Make the tests pass!

// I AM NOT DONE
//
// cleanupAll должна вернуть исходную ошибку работы, даже если очистка
// тоже запаниковала. Сейчас вторая паника «перекрывает» первую.
// Тренирует: новая паника в defer заменяет текущую.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

func cleanupAll(work, cleanup func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	defer func() {
		cleanup()
	}()
	work()
	return nil
}

func TestCleanupAll(t *testing.T) {
	err := cleanupAll(func() { panic("work failed") }, func() { panic("cleanup failed") })
	if err == nil || err.Error() != "work failed" {
		t.Errorf("err = %v, want work failed", err)
	}
}
