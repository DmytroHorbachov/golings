// functions42
// Make the tests pass!

// I AM NOT DONE
//
// NewServer принимает опции-функции, которые меняют настройки по умолчанию.
// Опции сейчас создаются, но не применяются, и одна из них пишет не в то поле.
// Тренирует: паттерн functional options.
// Сложность: medium
package main_test

import "testing"

type Server struct {
	Port    int
	Timeout int
}

type Option func(*Server)

func WithPort(p int) Option { return func(s *Server) { s.Port = p } }

func WithTimeout(t int) Option { return func(s *Server) { s.Port = t } }

func NewServer(opts ...Option) *Server {
	s := &Server{Port: 80, Timeout: 30}
	return s
}

func TestNewServer(t *testing.T) {
	s := NewServer(WithPort(8080), WithTimeout(5))
	if s.Port != 8080 || s.Timeout != 5 {
		t.Errorf("NewServer = %+v, want Port 8080 Timeout 5", *s)
	}
	d := NewServer()
	if d.Port != 80 || d.Timeout != 30 {
		t.Errorf("default server = %+v", *d)
	}
}
