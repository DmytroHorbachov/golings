// anonymous_functions_x004: Несколько разделителей
// Make the tests pass!
// I AM NOT DONE
//
// splitList делит строку по ';' и ','.
// Тренирует: литерал-предикат для strings.FieldsFunc.
// Сложность: easy
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func splitList(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == ';'
	})
}

func TestSplitList(t *testing.T) {
	if got := splitList("a,b;c;;d"); !reflect.DeepEqual(got, []string{"a", "b", "c", "d"}) {
		t.Errorf("splitList = %v", got)
	}
}
