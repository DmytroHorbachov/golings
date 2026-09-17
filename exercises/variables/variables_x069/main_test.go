// variables_x069: Анонимная структура
// Make the tests pass!
// I AM NOT DONE
//
// Функция serverConfig должна вернуть адрес вида "host:port".
// Настройки хранятся в переменной анонимного структурного типа.
// Тренирует: объявление переменной с анонимной структурой.
// Сложность: medium
package main_test

import (
	"strconv"
	"testing"
)

func serverAddr() string {
	cfg := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 80,
	}
	return cfg.Host + ":" + strconv.Itoa(cfg.Port)
}

func TestServerAddr(t *testing.T) {
	if got := serverAddr(); got != "127.0.0.1:8080" {
		t.Errorf("serverAddr() = %q, want 127.0.0.1:8080", got)
	}
}
