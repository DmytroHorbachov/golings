// slices97
// Make the tests pass!

// I AM NOT DONE
//
// addTask добавляет задачу в конец списка.
// Тренирует: append возвращает новый срез.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func addTask(tasks []string, t string) []string {
	_ = append(tasks, t)
	return tasks
}

func TestAddTask(t *testing.T) {
	got := addTask([]string{"wash"}, "cook")
	if !reflect.DeepEqual(got, []string{"wash", "cook"}) {
		t.Errorf("addTask = %v", got)
	}
}
