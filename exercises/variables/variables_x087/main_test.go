// variables_x087: Тайм-аут в миллисекундах
// Make the tests pass!
// I AM NOT DONE
//
// Константа Timeout должна быть равна 2.5 секунды, а функция —
// возвращать её в миллисекундах как int64.
// Тренирует: типизированные константы time.Duration и метод Milliseconds.
// Сложность: medium
package main_test

import (
	"testing"
	"time"
)

const Timeout = 2 * time.Second

func timeoutMs() int64 {
	return int64(Timeout.Seconds())
}

func TestTimeoutMs(t *testing.T) {
	if Timeout != 2500*time.Millisecond {
		t.Errorf("Timeout = %v, want 2.5s", Timeout)
	}
	if got := timeoutMs(); got != 2500 {
		t.Errorf("timeoutMs() = %d, want 2500", got)
	}
}
