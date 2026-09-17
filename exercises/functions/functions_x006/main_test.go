// functions_x006: Передать, а не вызвать
// Make the tests pass!
// I AM NOT DONE
//
// Функция apply ожидает функцию, а ей передают результат вызова.
// Код не компилируется.
// Тренирует: разницу между f и f().
// Сложность: easy
package main_test

import "testing"

func greet() string { return "hello" }

func apply(f func() string) string {
	return f() + "!"
}

func shout() string {
	return apply(greet())
}

func TestShout(t *testing.T) {
	if got := shout(); got != "hello!" {
		t.Errorf("shout() = %q, want %q", got, "hello!")
	}
}
