// arrays87
// Make the tests pass!

// I AM NOT DONE
//
// makeCounters должна вернуть три независимых счётчика.
// Все указатели указывают на одну и ту же переменную.
// Тренирует: копирование указателей не копирует данные.
// Сложность: hard
package main_test

import "testing"

func makeCounters() [3]*int {
	var out [3]*int
	n := 0
	for i := range out {
		out[i] = &n
	}
	return out
}

func TestMakeCounters(t *testing.T) {
	c := makeCounters()
	*c[0] += 1
	*c[1] += 10
	if *c[0] != 1 || *c[1] != 10 || *c[2] != 0 {
		t.Errorf("counters = %d %d %d", *c[0], *c[1], *c[2])
	}
}
