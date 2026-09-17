// structs47
// Make the tests pass!

// I AM NOT DONE
//
// NewServer возвращает сервер с портом 8080 и тайм-аутом 30.
// Тренирует: функции-конструкторы.
// Сложность: easy
package main_test

import "testing"

type Server struct {
	Port    int
	Timeout int
}

func NewServer() *Server {
	return &Server{Port: 8080}
}

func TestNewServer(t *testing.T) {
	s := NewServer()
	if s.Port != 8080 || s.Timeout != 30 {
		t.Errorf("NewServer = %+v", *s)
	}
}
