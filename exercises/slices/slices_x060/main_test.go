// slices_x060: Декартово произведение
// Make the tests pass!
// I AM NOT DONE
//
// product возвращает все пары (a, b) для a из первого среза и b из второго.
// Тренирует: вложенные циклы и срез результата с заранее известной ёмкостью.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func product(a []string, b []int) []string {
	out := make([]string, 0, len(a)*len(b))
	for i := range a {
		out = append(out, a[i]+string(rune('0'+b[i])))
	}
	return out
}

func TestProduct(t *testing.T) {
	got := product([]string{"a", "b"}, []int{1, 2, 3})
	if !reflect.DeepEqual(got, []string{"a1", "a2", "a3", "b1", "b2", "b3"}) {
		t.Errorf("product = %v", got)
	}
}
