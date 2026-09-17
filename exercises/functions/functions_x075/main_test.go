// functions_x075: Срез в ...interface{}
// Make the tests pass!
// I AM NOT DONE
//
// logf должна передать аргументы в fmt.Sprint так, чтобы они вывелись
// отдельными значениями, а не одним срезом.
// Тренирует: разницу между f(args) и f(args...) для ...interface{}.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

func logf(args ...interface{}) string {
	return fmt.Sprint(args)
}

func TestLogf(t *testing.T) {
	if got := logf("id=", 7); got != "id=7" {
		t.Errorf("logf(\"id=\", 7) = %q, want %q", got, "id=7")
	}
}
