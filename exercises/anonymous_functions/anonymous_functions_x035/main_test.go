// anonymous_functions_x035: Литерал с указателем
// Make the tests pass!
// I AM NOT DONE
//
// bumpAll увеличивает все счётчики с помощью литерала, принимающего *int.
// Тренирует: разыменование указателя внутри литерала.
// Сложность: easy
package main_test

import "testing"

func bumpAll(a, b *int) {
	inc := func(p *int) {
		p = nil
	}
	inc(a)
	inc(b)
}

func TestBumpAll(t *testing.T) {
	x, y := 1, 5
	bumpAll(&x, &y)
	if x != 2 || y != 6 {
		t.Errorf("x=%d y=%d", x, y)
	}
}
