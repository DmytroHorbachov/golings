// range55
// Make the tests pass!

// I AM NOT DONE
//
// lengthHistogram возвращает срез, где элемент i — число слов длины i (в символах).
// Тренирует: range по словам и расширение результата по необходимости.
// Сложность: medium
package main_test

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func lengthHistogram(text string) []int {
	var h []int
	for _, w := range strings.Fields(text) {
		n := len(w)
		h = append(h, n)
	}
	return h
}

func TestLengthHistogram(t *testing.T) {
	_ = utf8.RuneCountInString
	if got := lengthHistogram("я и ты мы go"); !reflect.DeepEqual(got, []int{0, 2, 3}) {
		t.Errorf("lengthHistogram = %v", got)
	}
}
