// functions19
// Make the tests pass!

// I AM NOT DONE
//
// Функция-обработчик обязана соответствовать типу Handler, но id ей не нужен.
// Код не компилируется: сигнатура не совпадает с типом.
// Тренирует: функциональные типы и пустой идентификатор в параметрах.
// Сложность: easy
package main_test

import "testing"

type Handler func(id int, name string) string

func run(h Handler) string {
	return h(42, "gopher")
}

func hello(name string) string {
	return "hello " + name
}

func TestRunHello(t *testing.T) {
	if got := run(hello); got != "hello gopher" {
		t.Errorf("run(hello) = %q, want %q", got, "hello gopher")
	}
}
