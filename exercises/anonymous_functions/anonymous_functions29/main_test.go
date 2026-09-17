// anonymous_functions29
// Make the tests pass!

// I AM NOT DONE
//
// Литерал должен считать события в исходной map статистики. После ротации
// переменная stats указывает на новую map, и литерал пишет уже туда.
// Тренирует: литерал видит новое значение захваченной переменной.
// Сложность: hard
package main_test

import "testing"

func run() (old, fresh map[string]int) {
	stats := map[string]int{}
	count := func(e string) { stats[e]++ }
	count("a")
	old = stats
	stats = map[string]int{}
	count("b")
	return old, stats
}

func TestRun(t *testing.T) {
	old, fresh := run()
	if old["a"] != 1 || old["b"] != 1 || len(fresh) != 0 {
		t.Errorf("old = %v, fresh = %v", old, fresh)
	}
}
