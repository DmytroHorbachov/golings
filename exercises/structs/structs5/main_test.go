// structs5
// Make the tests pass!

// I AM NOT DONE
//
// Код пытается изменить поле структуры, возвращённой функцией.
// Не компилируется: результат вызова не адресуем.
// Тренирует: присваивать можно только адресуемым значениям.
// Сложность: hard
package main_test

import "testing"

type Options struct {
	Retries int
	Verbose bool
}

func defaults() Options { return Options{Retries: 3} }

func verboseDefaults() Options {
	defaults().Verbose = true
	return defaults()
}

func TestVerboseDefaults(t *testing.T) {
	if o := verboseDefaults(); !o.Verbose || o.Retries != 3 {
		t.Errorf("verboseDefaults = %+v", o)
	}
}
