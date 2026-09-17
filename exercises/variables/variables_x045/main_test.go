// variables_x045: Типизированный nil
// Make the tests pass!
// I AM NOT DONE
//
// Функция validate возвращает error и должна вернуть nil, если ошибок нет.
// Однако результат никогда не равен nil.
// Тренирует: интерфейс, содержащий nil-указатель, сам не равен nil.
// Сложность: hard
package main_test

import "testing"

type ValidationError struct{ Field string }

func (e *ValidationError) Error() string { return "invalid " + e.Field }

func validate(name string) error {
	var verr *ValidationError
	if name == "" {
		verr = &ValidationError{Field: "name"}
	}
	return verr
}

func TestValidate(t *testing.T) {
	if err := validate("gopher"); err != nil {
		t.Errorf("validate(gopher) = %v, want nil", err)
	}
	if err := validate(""); err == nil {
		t.Errorf("validate(\"\") should return an error")
	}
}
