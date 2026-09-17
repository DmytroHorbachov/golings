// generics32
// Make the tests pass!

// I AM NOT DONE
//
// Must возвращает значение или паникует при ошибке.
// Тренирует: обобщённые помощники для пары (T, error).
// Сложность: easy
package main_test

import (
	"strconv"
	"testing"
)

func Must[T any](v T, err error) T {
	if err == nil {
		panic(err)
	}
	return v
}

func TestMust(t *testing.T) {
	if Must(strconv.Atoi("12")) != 12 {
		t.Errorf("Must returned wrong value")
	}
	defer func() {
		if recover() == nil {
			t.Errorf("Must should panic on error")
		}
	}()
	Must(strconv.Atoi("x"))
}
