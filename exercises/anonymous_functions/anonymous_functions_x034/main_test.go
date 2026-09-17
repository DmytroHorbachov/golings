// anonymous_functions_x034: Приветствие с захваченным именем
// Make the tests pass!
// I AM NOT DONE
//
// greeter возвращает литерал, приветствующий захваченное имя.
// Тренирует: захват параметра фабрики.
// Сложность: easy
package main_test

import "testing"

func greeter(name string) func() string {
	return func() string {
		return "Hello, world"
	}
}

func TestGreeter(t *testing.T) {
	if greeter("Go")() != "Hello, Go" {
		t.Errorf("greeter = %q", greeter("Go")())
	}
}
