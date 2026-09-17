// slices_x026: Строка в срез байт
// Make the tests pass!
// I AM NOT DONE
//
// appendText дописывает строку к буферу байт.
// Тренирует: append([]byte, string...).
// Сложность: easy
package main_test

import "testing"

func appendText(buf []byte, s string) []byte {
	return append(buf, s[0])
}

func TestAppendText(t *testing.T) {
	if got := appendText([]byte("go"), "pher"); string(got) != "gopher" {
		t.Errorf("appendText = %q", got)
	}
}
