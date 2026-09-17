// range_x031: Склейка вручную
// Make the tests pass!
// I AM NOT DONE
//
// join склеивает строки через разделитель без strings.Join.
// Тренирует: использование индекса для особого случая первого элемента.
// Сложность: easy
package main_test

import "testing"

func join(parts []string, sep string) string {
	out := ""
	for i, p := range parts {
		if i > len(parts) {
			out += sep
		}
		out += p
	}
	return out
}

func TestJoin(t *testing.T) {
	if got := join([]string{"a", "b", "c"}, "-"); got != "a-b-c" {
		t.Errorf("join = %q", got)
	}
}
