// arrays22
// Make the tests pass!

// I AM NOT DONE
//
// firstTwo должна вернуть срез с первыми двумя элементами массива.
// Тренирует: получение среза из массива.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func firstTwo(a *[4]string) []string {
	return a[1:3]
}

func TestFirstTwo(t *testing.T) {
	a := [4]string{"w", "x", "y", "z"}
	if got := firstTwo(&a); !reflect.DeepEqual(got, []string{"w", "x"}) {
		t.Errorf("firstTwo = %v", got)
	}
}
