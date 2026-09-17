// primitive_types35
// Make the tests pass!

// I AM NOT DONE
//
// cString превращает буфер фиксированного размера с C-строкой (оканчивается
// нулевым байтом) в строку Go. Сейчас в строку попадают нулевые байты.
// Тренирует: string(buf) копирует все байты, включая \x00.
// Сложность: hard
package main_test

import (
	"bytes"
	"testing"
)

func cString(buf []byte) string {
	return string(buf)
}

func TestCString(t *testing.T) {
	_ = bytes.IndexByte
	buf := make([]byte, 16)
	copy(buf, "gopher")
	if got := cString(buf); got != "gopher" {
		t.Errorf("cString = %q, want gopher", got)
	}
	if got := cString([]byte("full")); got != "full" {
		t.Errorf("cString(no terminator) = %q", got)
	}
}
