// range_x028: Первая ошибка
// Make the tests pass!
// I AM NOT DONE
//
// firstError возвращает первую ненулевую ошибку из среза.
// Тренирует: range по срезу интерфейсов.
// Сложность: easy
package main_test

import (
	"errors"
	"testing"
)

func firstError(errs []error) error {
	for _, err := range errs {
		if err == nil {
			return err
		}
	}
	return nil
}

func TestFirstError(t *testing.T) {
	e1 := errors.New("disk")
	if firstError([]error{nil, e1, errors.New("net")}) != e1 {
		t.Errorf("firstError should return disk error")
	}
	if firstError([]error{nil, nil}) != nil {
		t.Errorf("firstError should return nil")
	}
}
