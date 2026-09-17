// functions_x070: Переприсваивание map в функции
// Make the tests pass!
// I AM NOT DONE
//
// resetStats должна очистить статистику вызывающего кода.
// Внутри функции map переприсваивается, и снаружи ничего не меняется.
// Тренирует: map — ссылочный тип, но сама переменная передаётся по значению.
// Сложность: hard
package main_test

import "testing"

func resetStats(stats map[string]int) {
	stats = map[string]int{}
}

func TestResetStats(t *testing.T) {
	stats := map[string]int{"hits": 10, "misses": 2}
	resetStats(stats)
	if len(stats) != 0 {
		t.Errorf("after reset stats = %v, want empty", stats)
	}
}
