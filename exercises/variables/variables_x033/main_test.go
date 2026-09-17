// variables_x033: Аргументы defer
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна записать в журнал ИТОГОВОЕ значение счётчика при выходе.
// В журнал попадает начальное значение.
// Тренирует: аргументы отложенного вызова вычисляются в момент defer.
// Сложность: hard
package main_test

import "testing"

func process(items []string, log *[]int) {
	count := 0
	defer record(log, count)
	for range items {
		count++
	}
}

func record(log *[]int, v int) {
	*log = append(*log, v)
}

func TestProcess(t *testing.T) {
	var log []int
	process([]string{"a", "b", "c"}, &log)
	if len(log) != 1 || log[0] != 3 {
		t.Errorf("log = %v, want [3]", log)
	}
}
