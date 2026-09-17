// variables_x017: Тип time.Duration
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть длительность тайм-аута: seconds секунд.
// Сейчас умножаются значения разных типов.
// Тренирует: преобразование к именованному типу time.Duration.
// Сложность: easy
package main_test

import (
	"testing"
	"time"
)

func timeout(seconds int) time.Duration {
	return seconds * time.Second
}

func TestTimeout(t *testing.T) {
	if got := timeout(3); got != 3*time.Second {
		t.Errorf("timeout(3) = %v, want 3s", got)
	}
}
