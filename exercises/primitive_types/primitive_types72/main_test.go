// primitive_types72
// Make the tests pass!

// I AM NOT DONE
//
// elapsedMs возвращает разницу между двумя моментами в миллисекундах.
// Для интервала больше 25 дней результат становится отрицательным.
// Тренирует: int32 хранит лишь около 2.1 миллиарда (≈24.8 суток в мс).
// Сложность: hard
package main_test

import (
	"testing"
	"time"
)

func elapsedMs(from, to time.Time) int32 {
	return int32(to.Sub(from).Milliseconds())
}

func TestElapsedMs(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if got := elapsedMs(start, start.Add(1500*time.Millisecond)); got != 1500 {
		t.Errorf("elapsedMs(1.5s) = %d", got)
	}
	if got := elapsedMs(start, start.Add(30*24*time.Hour)); got != 2592000000 {
		t.Errorf("elapsedMs(30 days) = %d, want 2592000000", got)
	}
}
