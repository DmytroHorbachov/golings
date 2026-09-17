// anonymous_functions_x040: Все ошибки валидации
// Make the tests pass!
// I AM NOT DONE
//
// validate прогоняет значение через все проверки-литералы и собирает все ошибки.
// Тренирует: срез литералов и накопление результатов.
// Сложность: medium
package main_test

import (
	"errors"
	"strings"
	"testing"
)

var rules = []func(string) error{
	func(s string) error {
		if len(s) < 8 {
			return errors.New("too short")
		}
		return nil
	},
	func(s string) error {
		if !strings.ContainsAny(s, "0123456789") {
			return errors.New("no digit")
		}
		return nil
	},
}

func validate(pw string) []error {
	var errs []error
	for _, rule := range rules {
		err := rule(pw)
		if err != nil {
			return []error{err}
		}
	}
	return errs
}

func TestValidate(t *testing.T) {
	if errs := validate("abc"); len(errs) != 2 {
		t.Errorf("validate(abc) = %v, want 2 errors", errs)
	}
	if errs := validate("password1"); len(errs) != 0 {
		t.Errorf("validate(password1) = %v", errs)
	}
}
