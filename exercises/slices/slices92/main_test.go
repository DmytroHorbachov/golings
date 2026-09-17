// slices92
// Make the tests pass!

// I AM NOT DONE
//
// activeOnly фильтрует список, но вызывающий ожидает, что исходный срез
// останется прежним. Фильтрация через s[:0] его портит.
// Тренирует: out := s[:0] переиспользует массив исходного среза.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

type Task struct {
	Name   string
	Active bool
}

func activeOnly(tasks []Task) []Task {
	out := tasks[:0]
	for _, t := range tasks {
		if t.Active {
			out = append(out, t)
		}
	}
	return out
}

func TestActiveOnly(t *testing.T) {
	tasks := []Task{{"a", false}, {"b", true}, {"c", true}}
	got := activeOnly(tasks)
	if len(got) != 2 || got[0].Name != "b" {
		t.Errorf("activeOnly = %v", got)
	}
	if !reflect.DeepEqual(tasks, []Task{{"a", false}, {"b", true}, {"c", true}}) {
		t.Errorf("tasks modified: %v", tasks)
	}
}
