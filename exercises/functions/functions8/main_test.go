// functions8
// Make the tests pass!

// I AM NOT DONE
//
// readAll возвращает данные и закрывает файл в defer.
// Ошибка чтения не должна теряться, а ошибка закрытия должна возвращаться,
// только если других ошибок не было.
// Тренирует: изменение именованного результата в defer без затирания ошибки.
// Сложность: hard
package main_test

import (
	"errors"
	"testing"
)

type file struct {
	data     string
	closeErr error
}

func (f *file) Close() error { return f.closeErr }

func readAll(f *file) (s string, err error) {
	defer func() {
		err = f.Close()
	}()
	if f.data == "" {
		return "", errors.New("empty file")
	}
	return f.data, nil
}

func TestReadAll(t *testing.T) {
	if _, err := readAll(&file{}); err == nil || err.Error() != "empty file" {
		t.Errorf("readAll(empty) err = %v, want empty file", err)
	}
	closeErr := errors.New("close failed")
	if _, err := readAll(&file{data: "x", closeErr: closeErr}); err != closeErr {
		t.Errorf("readAll(close error) err = %v, want close failed", err)
	}
	if _, err := readAll(&file{data: "x", closeErr: closeErr}); err == nil {
		t.Errorf("close error must not be lost")
	}
	if s, err := readAll(&file{data: "ok"}); err != nil || s != "ok" {
		t.Errorf("readAll(ok) = %q, %v", s, err)
	}
}
