// range16
// Make the tests pass!

// I AM NOT DONE
//
// runeWidths возвращает ширину в байтах каждого символа, декодируя строку вручную.
// Цикл сдвигается на один байт и «видит» лишние символы.
// Тренирует: utf8.DecodeRuneInString возвращает ширину, на которую нужно сдвинуться.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func runeWidths(s string) []int {
	var out []int
	for i := 0; i < len(s); i++ {
		_, w := utf8.DecodeRuneInString(s[i:])
		out = append(out, w)
	}
	return out
}

func TestRuneWidths(t *testing.T) {
	if got := runeWidths("aλ€"); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("runeWidths = %v, want [1 2 3]", got)
	}
}
