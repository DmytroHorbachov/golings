// range75
// Make the tests pass!

// I AM NOT DONE
//
// vowelsPerWord возвращает количество гласных в каждом слове текста.
// Тренирует: вложенный range: по словам и по рунам слова.
// Сложность: medium
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func vowelsPerWord(text string) []int {
	var out []int
	n := 0
	for _, w := range strings.Fields(text) {
		for _, r := range w {
			if strings.ContainsRune("aeiouаеёиоуыэюя", r) {
				n++
			}
		}
	}
	out = append(out, n)
	return out
}

func TestVowelsPerWord(t *testing.T) {
	if got := vowelsPerWord("Gopher Ест Мёд"); !reflect.DeepEqual(got, []int{2, 1, 1}) {
		t.Errorf("vowelsPerWord = %v", got)
	}
}
