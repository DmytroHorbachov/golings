// structs7
// Make the tests pass!

// I AM NOT DONE
//
// Service встраивает Logger и использует его метод Log с префиксом имени сервиса.
// Тренирует: встраивание для переиспользования поведения.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type Logger struct {
	prefix string
	lines  []string
}

func (l *Logger) Log(msg string) { l.lines = append(l.lines, l.prefix+msg) }

type Service struct {
	log  Logger
	Name string
}

func NewService(name string) *Service {
	return &Service{Name: name}
}

func TestService(t *testing.T) {
	s := NewService("auth")
	s.Log("started")
	if !reflect.DeepEqual(s.lines, []string{"[auth] started"}) {
		t.Errorf("lines = %v", s.lines)
	}
}
