// anonymous_functions45
// Make the tests pass!

// I AM NOT DONE
//
// collector возвращает литерал add, который дописывает в переданный срез.
// Вызывающий код не видит добавленных элементов.
// Тренирует: литерал захватывает параметр — копию заголовка среза.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func collector(dst []string) func(string) {
	return func(s string) { dst = append(dst, s) }
}

func gather() []string {
	var names []string
	add := collector(names)
	add("a")
	add("b")
	return names
}

func TestGather(t *testing.T) {
	if got := gather(); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("gather = %v", got)
	}
}
