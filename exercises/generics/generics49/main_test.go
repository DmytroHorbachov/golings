// generics49
// Make the tests pass!

// I AM NOT DONE
//
// Ptr возвращает указатель на копию переданного значения.
// Тренирует: обобщённые вспомогательные функции.
// Сложность: easy
package main_test

import "testing"

func Ptr[T any](v T) *T {
	return nil
}

func TestPtr(t *testing.T) {
	p := Ptr(42)
	if p == nil || *p != 42 {
		t.Fatalf("Ptr(42) = %v", p)
	}
	s := Ptr("go")
	if *s != "go" {
		t.Errorf("Ptr(go) = %q", *s)
	}
}
