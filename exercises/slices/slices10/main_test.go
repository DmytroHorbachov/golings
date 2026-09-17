// slices10
// Make the tests pass!

// I AM NOT DONE
//
// firstLine возвращает первую строку из буфера чтения. Буфер потом переиспользуется,
// и возвращённая строка «портится».
// Тренирует: подсрез ссылается на тот же массив, что и буфер.
// Сложность: hard
package main_test

import (
	"bytes"
	"testing"
)

func firstLine(buf []byte) []byte {
	i := bytes.IndexByte(buf, '\n')
	if i < 0 {
		i = len(buf)
	}
	return buf[:i]
}

func TestFirstLine(t *testing.T) {
	buf := []byte("hello\nworld")
	line := firstLine(buf)
	copy(buf, "XXXXX")
	if string(line) != "hello" {
		t.Errorf("line = %q, want hello", line)
	}
}
