// maps28
// Make the tests pass!

// I AM NOT DONE
//
// increment передаёт счётчик из map в функцию, принимающую *int.
// Код не компилируется: адрес элемента map взять нельзя.
// Тренирует: элементы map могут перемещаться при росте, поэтому не адресуемы.
// Сложность: hard
package main_test

import "testing"

func bump(p *int) { *p += 10 }

func increment(m map[string]int, k string) {
	bump(&m[k])
}

func TestIncrement(t *testing.T) {
	m := map[string]int{"a": 1}
	increment(m, "a")
	increment(m, "b")
	if m["a"] != 11 || m["b"] != 10 {
		t.Errorf("m = %v", m)
	}
}
