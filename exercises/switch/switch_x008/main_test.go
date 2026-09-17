// switch_x008: MIME-тип
// Make the tests pass!
// I AM NOT DONE
//
// mimeType возвращает MIME-тип по расширению файла.
// Тренирует: switch с инициализирующей инструкцией.
// Сложность: easy
package main_test

import (
	"path/filepath"
	"testing"
)

func mimeType(name string) string {
	switch ext := filepath.Ext(name); ext {
	case ".html":
		return "text/html"
	case "json":
		return "application/json"
	default:
		return "application/octet-stream"
	}
}

func TestMimeType(t *testing.T) {
	cases := map[string]string{"index.html": "text/html", "data.json": "application/json", "a.bin": "application/octet-stream"}
	for in, want := range cases {
		if got := mimeType(in); got != want {
			t.Errorf("mimeType(%s) = %s, want %s", in, got, want)
		}
	}
}
