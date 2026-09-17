// anonymous_functions102
// Make the tests pass!

// I AM NOT DONE
//
// makeReader возвращает литерал, читающий буфер. Отложенная очистка
// обнуляет захваченную переменную при выходе, и литерал видит пустую строку.
// Тренирует: defer выполняется до того, как вызывающий код вызовет литерал.
// Сложность: hard
package main_test

import "testing"

func makeReader(load func() string) func() string {
	buf := load()
	defer func() { buf = "" }()
	return func() string { return buf }
}

func TestMakeReader(t *testing.T) {
	read := makeReader(func() string { return "payload" })
	if got := read(); got != "payload" {
		t.Errorf("read = %q", got)
	}
}
