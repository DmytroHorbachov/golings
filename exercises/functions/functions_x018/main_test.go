// functions_x018: Разделители как функция
// Make the tests pass!
// I AM NOT DONE
//
// splitWords должна делить строку по любым небуквенным символам.
// Тренирует: strings.FieldsFunc и функции-предикаты.
// Сложность: easy
package main_test

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

func splitWords(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsLetter(r)
	})
}

func TestSplitWords(t *testing.T) {
	got := splitWords("go,is;fun!  yes")
	if want := []string{"go", "is", "fun", "yes"}; !reflect.DeepEqual(got, want) {
		t.Errorf("splitWords = %v, want %v", got, want)
	}
}
