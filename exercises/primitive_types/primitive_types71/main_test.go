// primitive_types71
// Make the tests pass!

// I AM NOT DONE
//
// fromBytes должна превратить срез байт в строку.
// Тренирует: преобразование []byte в string.
// Сложность: easy
package main_test

import (
	"fmt"
	"testing"
)

func fromBytes(b []byte) string {
	return fmt.Sprint(b)
}

func TestFromBytes(t *testing.T) {
	_ = fmt.Sprint
	if got := fromBytes([]byte{'g', 'o'}); got != "go" {
		t.Errorf("fromBytes = %q, want go", got)
	}
}
