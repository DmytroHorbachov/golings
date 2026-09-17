// variables_x046: Порядок инициализации
// Make the tests pass!
// I AM NOT DONE
//
// Переменная пакета greeting должна содержать "Hello, Gopher".
// Имя задаётся в init(), но greeting вычисляется раньше.
// Тренирует: порядок инициализации — сначала переменные пакета, потом init.
// Сложность: hard
package main_test

import "testing"

var name string

func init() {
	name = "Gopher"
}

var greeting = makeGreeting()

func makeGreeting() string {
	return "Hello, " + name
}

func TestGreeting(t *testing.T) {
	if greeting != "Hello, Gopher" {
		t.Errorf("greeting = %q, want %q", greeting, "Hello, Gopher")
	}
}
