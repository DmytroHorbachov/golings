// maps94
// Make the tests pass!

// I AM NOT DONE
//
// activeCount удаляет истёкшие сессии и возвращает количество оставшихся.
// Тренирует: len map после delete.
// Сложность: easy
package main_test

import "testing"

func activeCount(sessions map[string]int, now int) int {
	for id, expires := range sessions {
		if expires > now {
			delete(sessions, id)
		}
	}
	return len(sessions)
}

func TestActiveCount(t *testing.T) {
	s := map[string]int{"a": 5, "b": 15, "c": 20}
	if got := activeCount(s, 10); got != 2 {
		t.Errorf("activeCount = %d, want 2", got)
	}
}
