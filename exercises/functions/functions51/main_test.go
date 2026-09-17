// functions51
// Make the tests pass!

// I AM NOT DONE
//
// Функция steps должна вернуть журнал: "start", "work", "cleanup 2", "cleanup 1".
// Тренирует: отложенные вызовы выполняются в обратном порядке (LIFO).
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func steps() (log []string) {
	add := func(s string) { log = append(log, s) }
	add("start")
	defer add("cleanup 2")
	defer add("cleanup 1")
	add("work")
	return
}

func TestSteps(t *testing.T) {
	want := []string{"start", "work", "cleanup 2", "cleanup 1"}
	if got := steps(); !reflect.DeepEqual(got, want) {
		t.Errorf("steps() = %v, want %v", got, want)
	}
}
