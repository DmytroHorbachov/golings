// primitive_types_x087: Цикл по float
// Make the tests pass!
// I AM NOT DONE
//
// steps считает, сколько шагов по 0.1 нужно, чтобы дойти от 0 до 1.
// Условие != 1.0 никогда не выполняется точно, и цикл упирается в предохранитель.
// Тренирует: нельзя управлять циклом точным сравнением float.
// Сложность: hard
package main_test

import "testing"

func steps() int {
	n := 0
	for x := 0.0; x != 1.0 && n < 1000; x += 0.1 {
		n++
	}
	return n
}

func TestSteps(t *testing.T) {
	if got := steps(); got != 10 {
		t.Errorf("steps() = %d, want 10", got)
	}
}
