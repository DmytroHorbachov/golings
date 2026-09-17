// functions30
// Make the tests pass!

// I AM NOT DONE
//
// run превращает любую панику в ошибку. Паника может нести error или строку.
// Со строкой функция сама паникует внутри defer.
// Тренирует: recover возвращает interface{}; тип значения нужно проверять.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

func run(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	f()
	return nil
}

func TestRun(t *testing.T) {
	diskErr := fmt.Errorf("disk full")
	if err := run(func() { panic(diskErr) }); err != diskErr {
		t.Errorf("run(error panic) = %v, want disk full", err)
	}
	if err := run(func() { panic("oops") }); err == nil || err.Error() != "panic: oops" {
		t.Errorf("run(string panic) = %v, want panic: oops", err)
	}
}
