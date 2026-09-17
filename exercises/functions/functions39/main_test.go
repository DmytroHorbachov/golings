// functions39
// Make the tests pass!

// I AM NOT DONE
//
// Тип GreeterFunc — адаптер, позволяющий использовать функцию как Greeter.
// Сейчас у типа нет метода, и функция не удовлетворяет интерфейсу.
// Тренирует: методы у функциональных типов (как http.HandlerFunc).
// Сложность: medium
package main_test

import "testing"

type Greeter interface {
	Greet(name string) string
}

type GreeterFunc func(string) string

func welcome(g Greeter) string {
	return g.Greet("gopher")
}

func polite(name string) string { return "Good day, " + name }

func run() string {
	return welcome(polite)
}

func TestRun(t *testing.T) {
	if got := run(); got != "Good day, gopher" {
		t.Errorf("run() = %q", got)
	}
}
