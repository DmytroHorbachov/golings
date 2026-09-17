// generics98
// Make the tests pass!

// I AM NOT DONE
//
// IsNil должна возвращать true и для nil-указателя любого типа.
// any(v) == nil не срабатывает для nil-указателя.
// Тренирует: интерфейс с типизированным nil не равен nil.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func IsNil[T any](v T) bool {
	return any(v) == nil
}

func TestIsNil(t *testing.T) {
	_ = reflect.ValueOf
	var p *int
	var e error
	if !IsNil(p) || !IsNil(e) || IsNil(5) || IsNil(new(int)) {
		t.Errorf("IsNil works incorrectly")
	}
}
