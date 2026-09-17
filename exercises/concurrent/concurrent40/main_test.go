// concurrent40
// Make the tests pass!

// I AM NOT DONE
//
// lines передаёт строки через канал, но канал объявлен с другим типом элементов.
// Код не компилируется.
// Тренирует: типизированные каналы.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func lines(text string) []string {
	ch := make(chan []byte, 10)
	for _, l := range strings.Split(text, "\n") {
		ch <- l
	}
	close(ch)
	var out []string
	for l := range ch {
		out = append(out, l)
	}
	return out
}

func TestLines(t *testing.T) {
	if got := lines("a\nb"); len(got) != 2 || got[1] != "b" {
		t.Errorf("lines = %v", got)
	}
}
