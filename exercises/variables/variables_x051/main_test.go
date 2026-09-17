// variables_x051: Дробные секунды
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна превратить дробное число секунд (например 1.5) в time.Duration.
// Сейчас дробная часть теряется.
// Тренирует: time.Duration — это целое число наносекунд.
// Сложность: hard
package main_test

import (
	"testing"
	"time"
)

func fromSeconds(secs float64) time.Duration {
	return time.Duration(secs) * time.Second
}

func TestFromSeconds(t *testing.T) {
	if got := fromSeconds(1.5); got != 1500*time.Millisecond {
		t.Errorf("fromSeconds(1.5) = %v, want 1.5s", got)
	}
	if got := fromSeconds(0.25); got != 250*time.Millisecond {
		t.Errorf("fromSeconds(0.25) = %v, want 250ms", got)
	}
}
