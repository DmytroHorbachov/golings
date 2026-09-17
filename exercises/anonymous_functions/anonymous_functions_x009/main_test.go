// anonymous_functions_x009: Префиксатор
// Make the tests pass!
// I AM NOT DONE
//
// withPrefix возвращает функцию, добавляющую префикс к строке.
// Тренирует: захват параметра в замыкании.
// Сложность: easy
package main_test

import "testing"

func withPrefix(p string) func(string) string {
	return func(s string) string {
		return s + p
	}
}

func TestWithPrefix(t *testing.T) {
	warn := withPrefix("[warn] ")
	if warn("disk") != "[warn] disk" {
		t.Errorf("warn = %q", warn("disk"))
	}
}
