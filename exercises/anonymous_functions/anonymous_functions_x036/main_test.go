// anonymous_functions_x036: Итератор по срезу
// Make the tests pass!
// I AM NOT DONE
//
// iterator возвращает литерал, который при каждом вызове выдаёт следующий элемент
// и false, когда элементы закончились.
// Тренирует: замыкание-итератор с позицией.
// Сложность: medium
package main_test

import "testing"

func iterator(items []string) func() (string, bool) {
	pos := 0
	return func() (string, bool) {
		pos++
		return items[pos], true
	}
}

func TestIterator(t *testing.T) {
	next := iterator([]string{"a", "b"})
	var got []string
	for v, ok := next(); ok; v, ok = next() {
		got = append(got, v)
	}
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Errorf("iterated %v", got)
	}
}
