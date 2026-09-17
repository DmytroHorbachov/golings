// functions93
// Make the tests pass!

// I AM NOT DONE
//
// validate прогоняет значение через список проверок и возвращает первую ошибку.
// Сейчас возвращается последняя ошибка или nil.
// Тренирует: срез функций и ранний выход.
// Сложность: medium
package main_test

import (
	"errors"
	"testing"
)

type check func(string) error

func notEmpty(s string) error {
	if s == "" {
		return errors.New("empty")
	}
	return nil
}

func maxLen(n int) check {
	return func(s string) error {
		if len(s) > n {
			return errors.New("too long")
		}
		return nil
	}
}

func validate(s string, checks ...check) error {
	var err error
	for _, c := range checks {
		err = c(s)
	}
	return err
}

func TestValidate(t *testing.T) {
	if err := validate("", notEmpty, maxLen(3)); err == nil || err.Error() != "empty" {
		t.Errorf("validate(\"\") = %v, want empty", err)
	}
	if err := validate("golang", maxLen(3), notEmpty); err == nil || err.Error() != "too long" {
		t.Errorf("validate(golang) = %v, want too long", err)
	}
	if err := validate("go", notEmpty, maxLen(3)); err != nil {
		t.Errorf("validate(go) = %v, want nil", err)
	}
}
