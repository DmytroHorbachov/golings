// slices80
// Make the tests pass!

// I AM NOT DONE
//
// variants строит два пути, продолжающих общий префикс разными шагами.
// Второй вариант затирает первый.
// Тренирует: два append к срезу с запасом ёмкости используют один массив.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func extend(base []string, step string) []string {
	return append(base, step)
}

func variants(base []string) ([]string, []string) {
	return extend(base, "left"), extend(base, "right")
}

func TestVariants(t *testing.T) {
	base := make([]string, 0, 10)
	base = append(base, "start")
	l, r := variants(base)
	if !reflect.DeepEqual(l, []string{"start", "left"}) || !reflect.DeepEqual(r, []string{"start", "right"}) {
		t.Errorf("variants = %v, %v", l, r)
	}
}
