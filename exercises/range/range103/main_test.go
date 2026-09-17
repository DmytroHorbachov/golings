// range103
// Make the tests pass!

// I AM NOT DONE
//
// resetAll очищает счётчики во всех map среза, заменяя их новыми пустыми map.
// Присваивание переменной цикла не меняет элементы среза.
// Тренирует: v в range — копия ссылки на map; её переприсваивание не видно снаружи.
// Сложность: hard
package main_test

import "testing"

func resetAll(stats []map[string]int) {
	for _, m := range stats {
		m = map[string]int{}
		_ = m
	}
}

func TestResetAll(t *testing.T) {
	a := map[string]int{"x": 1}
	stats := []map[string]int{a, {"y": 2}}
	resetAll(stats)
	if len(stats[0]) != 0 || len(stats[1]) != 0 {
		t.Errorf("stats = %v", stats)
	}
	if a["x"] != 1 {
		t.Errorf("old map should stay untouched: %v", a)
	}
}
