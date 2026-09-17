// functions74
// Make the tests pass!

// I AM NOT DONE
//
// joinAll должна передать все строки в fmt.Sprint как отдельные аргументы.
// Код не компилируется.
// Тренирует: []T нельзя передать как ...interface{}, даже если T подходит.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

func joinAll(names []string) string {
	return fmt.Sprint(names...)
}

func TestJoinAll(t *testing.T) {
	if got := joinAll([]string{"go", "-", "pher"}); got != "go-pher" {
		t.Errorf("joinAll = %q, want %q", got, "go-pher")
	}
}
