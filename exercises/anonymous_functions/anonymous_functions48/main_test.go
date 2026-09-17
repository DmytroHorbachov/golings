// anonymous_functions48
// Make the tests pass!

// I AM NOT DONE
//
// process помечает задачу выполненной в отложенном литерале.
// Тренирует: defer func() { ... }().
// Сложность: easy
package main_test

import "testing"

type Task struct{ Done bool }

func process(t *Task) {
	defer func() {
		t.Done = false
	}()
}

func TestProcess(t *testing.T) {
	var task Task
	process(&task)
	if !task.Done {
		t.Errorf("task not done")
	}
}
