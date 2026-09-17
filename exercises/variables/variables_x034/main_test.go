// variables_x034: Снимок значения
// Make the tests pass!
// I AM NOT DONE
//
// Функция makeGreeter должна «запомнить» имя в момент создания приветствия.
// Последующее изменение переменной name не должно влиять на результат.
// Тренирует: замыкания захватывают переменную, а не её значение.
// Сложность: hard
package main_test

import "testing"

func makeGreeters() (func() string, func() string) {
	name := "Alice"
	first := func() string { return "Hi, " + name }
	name = "Bob"
	second := func() string { return "Hi, " + name }
	return first, second
}

func TestGreeters(t *testing.T) {
	first, second := makeGreeters()
	if got := first(); got != "Hi, Alice" {
		t.Errorf("first() = %q, want %q", got, "Hi, Alice")
	}
	if got := second(); got != "Hi, Bob" {
		t.Errorf("second() = %q, want %q", got, "Hi, Bob")
	}
}
