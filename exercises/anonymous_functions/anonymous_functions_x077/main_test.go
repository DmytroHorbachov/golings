// anonymous_functions_x077: Литерал не той сигнатуры
// Make the tests pass!
// I AM NOT DONE
//
// bytes.IndexFunc ждёт func(rune) bool. Литерал объявлен с параметром byte,
// и код не компилируется.
// Тренирует: типы параметров литерала должны совпадать точно.
// Сложность: hard
package main_test

import (
	"bytes"
	"testing"
)

func firstSpace(b []byte) int {
	return bytes.IndexFunc(b, func(c byte) bool { return c == ' ' })
}

func TestFirstSpace(t *testing.T) {
	if firstSpace([]byte("go lang")) != 2 {
		t.Errorf("firstSpace = %d", firstSpace([]byte("go lang")))
	}
}
