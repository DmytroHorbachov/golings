// slices76
// Make the tests pass!

// I AM NOT DONE
//
// preview возвращает не больше n первых комментариев. Если комментариев меньше,
// функция паникует.
// Тренирует: s[:n] при n > len(s) (и n > cap) — паника.
// Сложность: hard
package main_test

import "testing"

func preview(comments []string, n int) []string {
	return comments[:n]
}

func TestPreview(t *testing.T) {
	if got := preview([]string{"a", "b", "c"}, 2); len(got) != 2 {
		t.Errorf("preview(3, 2) = %v", got)
	}
	if got := preview([]string{"a"}, 5); len(got) != 1 {
		t.Errorf("preview(1, 5) = %v", got)
	}
}
