// anonymous_functions103
// Make the tests pass!

// I AM NOT DONE
//
// Logger вычисляет дорогое сообщение, только если уровень включён:
// сообщение передаётся литералом.
// Тренирует: отложенное вычисление аргументов через функцию.
// Сложность: medium
package main_test

import "testing"

type Logger struct {
	debug bool
	out   []string
}

func (l *Logger) Debug(msg string) {
	if l.debug {
		l.out = append(l.out, msg)
	}
}

func TestLazyDebug(t *testing.T) {
	calls := 0
	expensive := func() string { calls++; return "state dump" }
	l := &Logger{}
	l.Debug(expensive)
	l.debug = true
	l.Debug(expensive)
	if calls != 1 || len(l.out) != 1 {
		t.Errorf("calls = %d, out = %v", calls, l.out)
	}
}
