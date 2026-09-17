// structs_x035: Конструктор с проверкой
// Make the tests pass!
// I AM NOT DONE
//
// NewEmail создаёт адрес, проверяя наличие '@' и непустые части.
// Тренирует: конструктор, возвращающий (*T, error).
// Сложность: medium
package main_test

import (
	"errors"
	"strings"
	"testing"
)

type Email struct{ User, Domain string }

func NewEmail(s string) (*Email, error) {
	parts := strings.Split(s, "@")
	return &Email{parts[0], parts[1]}, nil
}

func TestNewEmail(t *testing.T) {
	_ = errors.New
	if e, err := NewEmail("ann@x.io"); err != nil || e.Domain != "x.io" {
		t.Errorf("NewEmail = %v, %v", e, err)
	}
	for _, bad := range []string{"annx.io", "@x.io", "a@", "a@b@c"} {
		if e, err := NewEmail(bad); err == nil || e != nil {
			t.Errorf("NewEmail(%q) should fail", bad)
		}
	}
}
