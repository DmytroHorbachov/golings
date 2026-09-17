// anonymous_functions_x025: Обработчики событий
// Make the tests pass!
// I AM NOT DONE
//
// on регистрирует литерал-обработчик события; emit вызывает его.
// Тренирует: хранение литералов в map.
// Сложность: easy
package main_test

import "testing"

func register(handlers map[string]func() string) {
	on := func(name string, h func() string) { handlers[name] = h }
	on("start", func() string { return "stopped" })
}

func TestHandlers(t *testing.T) {
	handlers := map[string]func() string{}
	register(handlers)
	if handlers["start"]() != "started" {
		t.Errorf("start handler = %q", handlers["start"]())
	}
}
