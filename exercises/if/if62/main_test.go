// if62
// Make the tests pass!

// I AM NOT DONE
//
// isEmpty должна вернуть true и для nil-среза, и для пустого среза.
// Тренирует: проверку длины вместо сравнения с nil.
// Сложность: easy
package main_test

import "testing"

func isEmpty(items []string) bool {
	if items == nil {
		return true
	}
	return false
}

func TestIsEmpty(t *testing.T) {
	if !isEmpty(nil) || !isEmpty([]string{}) {
		t.Errorf("nil and empty slices should be empty")
	}
	if isEmpty([]string{"x"}) {
		t.Errorf("[x] is not empty")
	}
}
