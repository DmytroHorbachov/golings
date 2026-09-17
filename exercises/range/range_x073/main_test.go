// range_x073: Затенение внешней переменной
// Make the tests pass!
// I AM NOT DONE
//
// lastNonEmpty должна вернуть последнюю непустую строку.
// Функция всегда возвращает пустую строку.
// Тренирует: := в заголовке range объявляет новые переменные.
// Сложность: hard
package main_test

import "testing"

func lastNonEmpty(lines []string) string {
	var last string
	for _, last := range lines {
		if last == "" {
			continue
		}
	}
	return last
}

func TestLastNonEmpty(t *testing.T) {
	if got := lastNonEmpty([]string{"a", "b", ""}); got != "b" {
		t.Errorf("lastNonEmpty = %q, want b", got)
	}
}
