// switch_x047: Счётчик классов символов
// Make the tests pass!
// I AM NOT DONE
//
// classify считает гласные, согласные, цифры и прочие символы ASCII-строки.
// Тренирует: switch без тега внутри цикла.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func classify(s string) (vowels, consonants, digits, other int) {
	for _, r := range strings.ToLower(s) {
		switch {
		case r >= 'a' && r <= 'z':
			consonants++
		case strings.ContainsRune("aeiou", r):
			vowels++
		default:
			other++
		}
	}
	return
}

func TestClassify(t *testing.T) {
	v, c, d, o := classify("Go 1.22 rocks!")
	if v != 2 || c != 5 || d != 3 || o != 4 {
		t.Errorf("classify = %d %d %d %d, want 2 5 3 4", v, c, d, o)
	}
}
