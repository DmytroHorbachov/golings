// anonymous_functions51
// Make the tests pass!

// I AM NOT DONE
//
// handler создаётся с захватом переменной db, а затем переменная переприсваивается nil
// при «закрытии». Обработчик должен использовать соединение, взятое при создании.
// Тренирует: литерал захватывает переменную, а не текущее значение указателя.
// Сложность: hard
package main_test

import "testing"

type DB struct{ Name string }

func setup() func() string {
	db := &DB{Name: "main"}
	handler := func() string { return db.Name }
	db = nil
	return handler
}

func TestSetup(t *testing.T) {
	if got := setup()(); got != "main" {
		t.Errorf("handler = %q", got)
	}
}
