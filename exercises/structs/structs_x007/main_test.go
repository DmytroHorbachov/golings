// structs_x007: Анонимная структура
// Make the tests pass!
// I AM NOT DONE
//
// config возвращает хост и порт из анонимной структуры.
// Тренирует: анонимные структурные типы.
// Сложность: easy
package main_test

import "testing"

func config() (string, int) {
	c := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 80,
	}
	return c.Host, c.Port
}

func TestConfig(t *testing.T) {
	if h, p := config(); h != "localhost" || p != 5432 {
		t.Errorf("config = %s:%d", h, p)
	}
}
