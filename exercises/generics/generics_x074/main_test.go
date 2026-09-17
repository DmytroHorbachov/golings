// generics_x074: type switch по T
// Make the tests pass!
// I AM NOT DONE
//
// Describe возвращает описание значения в зависимости от его типа.
// Код не компилируется: type switch нельзя применить к значению параметра типа.
// Тренирует: для type switch значение нужно привести к interface{}.
// Сложность: hard
package main_test

import (
	"strconv"
	"testing"
)

func Describe[T any](v T) string {
	switch x := v.(type) {
	case int:
		return "int " + strconv.Itoa(x)
	case string:
		return "string " + x
	}
	return "other"
}

func TestDescribe(t *testing.T) {
	if Describe(5) != "int 5" || Describe("go") != "string go" || Describe(1.5) != "other" {
		t.Errorf("Describe works incorrectly")
	}
}
