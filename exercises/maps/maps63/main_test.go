// maps63
// Make the tests pass!

// I AM NOT DONE
//
// countLines считает одинаковые строки из [][]byte. Код не компилируется:
// []byte нельзя использовать как ключ map.
// Тренирует: ключ map должен быть сравнимым типом.
// Сложность: hard
package main_test

import (
	"bytes"
	"testing"
)

func countLines(data []byte) map[string]int {
	m := map[string]int{}
	for _, line := range bytes.Split(data, []byte("\n")) {
		m[line]++
	}
	return m
}

func TestCountLines(t *testing.T) {
	m := countLines([]byte("a\nb\na"))
	if m["a"] != 2 || m["b"] != 1 {
		t.Errorf("countLines = %v", m)
	}
}
