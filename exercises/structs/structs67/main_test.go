// structs67
// Make the tests pass!

// I AM NOT DONE
//
// Handler встраивает *Logger. Handler{} создан без логгера, и вызов
// продвинутого метода паникует.
// Тренирует: встраивание указателя — это поле, которое может быть nil.
// Сложность: hard
package main_test

import "testing"

type Logger struct{ lines []string }

func (l *Logger) Log(s string) {
	l.lines = append(l.lines, s)
}

type Handler struct {
	*Logger
	Name string
}

func (h Handler) Serve() string {
	h.Log("serving " + h.Name)
	return "ok"
}

func TestHandler(t *testing.T) {
	if (Handler{Name: "quiet"}).Serve() != "ok" {
		t.Errorf("Serve failed")
	}
	h := Handler{Logger: &Logger{}, Name: "loud"}
	h.Serve()
	if len(h.lines) != 1 {
		t.Errorf("lines = %v", h.lines)
	}
}
