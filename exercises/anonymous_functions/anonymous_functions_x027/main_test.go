// anonymous_functions_x027: Ленивое значение по умолчанию
// Make the tests pass!
// I AM NOT DONE
//
// getOr возвращает значение или вызывает литерал, вычисляющий значение по умолчанию.
// Сейчас литерал вызывается всегда, даже когда значение есть.
// Тренирует: передачу функции для отложенного вычисления.
// Сложность: medium
package main_test

import "testing"

func getOr(m map[string]int, k string, def func() int) int {
	fallback := def()
	if v, ok := m[k]; ok {
		return v
	}
	return fallback
}

func TestGetOr(t *testing.T) {
	calls := 0
	def := func() int { calls++; return 42 }
	m := map[string]int{"a": 1}
	if getOr(m, "a", def) != 1 || getOr(m, "b", def) != 42 {
		t.Errorf("getOr works incorrectly")
	}
	if calls != 1 {
		t.Errorf("default computed %d times", calls)
	}
}
