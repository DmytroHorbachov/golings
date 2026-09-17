// anonymous_functions_x091: Утверждение к именованному типу функции
// Make the tests pass!
// I AM NOT DONE
//
// Обработчики хранятся как interface{}. Литерал имеет безымянный тип
// func(string) string, поэтому утверждение к Handler не срабатывает.
// Тренирует: динамический тип литерала — безымянный функциональный тип.
// Сложность: hard
package main_test

import "testing"

type Handler func(string) string

func call(h interface{}, arg string) string {
	if f, ok := h.(Handler); ok {
		return f(arg)
	}
	return "not a handler"
}

func TestCall(t *testing.T) {
	var h interface{} = func(s string) string { return "hi " + s }
	if got := call(h, "go"); got != "hi go" {
		t.Errorf("call = %q", got)
	}
	if got := call(Handler(func(s string) string { return s }), "x"); got != "x" {
		t.Errorf("call(Handler) = %q", got)
	}
}
