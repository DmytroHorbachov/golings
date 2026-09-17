// anonymous_functions14
// Make the tests pass!

// I AM NOT DONE
//
// transformAll применяет колбэк к каждой строке, но строки не меняются:
// результат колбэка отбрасывается.
// Тренирует: литерал, возвращающий значение, не меняет аргумент.
// Сложность: hard
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func transformAll(items []string, f func(string) string) {
	for i := range items {
		f(items[i])
	}
}

func TestTransformAll(t *testing.T) {
	s := []string{"a", "b"}
	transformAll(s, func(x string) string { return strings.ToUpper(x) + "!" })
	if !reflect.DeepEqual(s, []string{"A!", "B!"}) {
		t.Errorf("transformAll = %v", s)
	}
}
