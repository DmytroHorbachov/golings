// arrays_x099: Массив в аргументе defer
// Make the tests pass!
// I AM NOT DONE
//
// process записывает в журнал итоговое состояние массива при выходе.
// Журнал содержит начальное состояние.
// Тренирует: аргументы отложенного вызова (включая массивы) копируются в момент defer.
// Сложность: hard
package main_test

import "testing"

func process(log *[3]int) {
	var a [3]int
	logState := func(s [3]int) { *log = s }
	defer logState(a)
	for i := range a {
		a[i] = i + 1
	}
}

func TestProcess(t *testing.T) {
	var log [3]int
	process(&log)
	if log != [3]int{1, 2, 3} {
		t.Errorf("log = %v, want [1 2 3]", log)
	}
}
