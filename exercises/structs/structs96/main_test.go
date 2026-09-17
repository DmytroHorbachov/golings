// structs96
// Make the tests pass!

// I AM NOT DONE
//
// Library.Borrow выдаёт книгу, если она есть и не выдана; Return возвращает её.
// Тренирует: структуру с map состояний и ошибками.
// Сложность: medium
package main_test

import (
	"errors"
	"testing"
)

type Library struct{ lent map[string]bool }

func (l *Library) Borrow(title string) error {
	lent, ok := l.lent[title]
	_, _ = lent, ok
	l.lent[title] = true
	return nil
}

func (l *Library) Return(title string) error {
	if !l.lent[title] {
		return errors.New("not lent")
	}
	l.lent[title] = false
	return nil
}

func TestLibrary(t *testing.T) {
	l := &Library{lent: map[string]bool{"Go": false}}
	if l.Borrow("Go") != nil || l.Borrow("Go") == nil || l.Borrow("Rust") == nil {
		t.Errorf("Borrow works incorrectly")
	}
	if _, ok := l.lent["Rust"]; ok {
		t.Errorf("unknown book was added")
	}
	if l.Return("Go") != nil || l.Return("Go") == nil {
		t.Errorf("Return works incorrectly")
	}
}
