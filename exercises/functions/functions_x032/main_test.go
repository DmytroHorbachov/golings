// functions_x032: Успех — это nil
// Make the tests pass!
// I AM NOT DONE
//
// При успешной проверке validateAge должна возвращать nil.
// Тренирует: соглашение о возврате ошибок в Go.
// Сложность: easy
package main_test

import (
	"errors"
	"testing"
)

func validateAge(age int) error {
	if age < 0 || age > 150 {
		return errors.New("age out of range")
	}
	return errors.New("")
}

func TestValidateAge(t *testing.T) {
	if err := validateAge(30); err != nil {
		t.Errorf("validateAge(30) = %v, want nil", err)
	}
	if err := validateAge(-1); err == nil {
		t.Errorf("validateAge(-1) should fail")
	}
}
