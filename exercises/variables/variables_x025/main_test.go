// variables_x025: Функция new
// Make the tests pass!
// I AM NOT DONE
//
// Функция newCounter должна вернуть указатель на int со значением start.
// Тренирует: встроенную функцию new и разыменование указателя.
// Сложность: medium
package main_test

import "testing"

func newCounter(start int) *int {
	var p *int
	p = &start
	start = 0
	return p
}

func TestNewCounter(t *testing.T) {
	p := newCounter(7)
	if p == nil || *p != 7 {
		t.Fatalf("newCounter(7) should point to 7")
	}
	q := newCounter(7)
	*q = 8
	if *p != 7 {
		t.Errorf("counters must be independent, got %d", *p)
	}
}
