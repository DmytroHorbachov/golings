// anonymous_functions46
// Make the tests pass!

// I AM NOT DONE
//
// toggler возвращает замыкание, которое при каждом вызове меняет состояние на противоположное.
// Тренирует: замыкание над bool.
// Сложность: easy
package main_test

import "testing"

func toggler() func() bool {
	on := false
	return func() bool {
		on = true
		return on
	}
}

func TestToggler(t *testing.T) {
	tg := toggler()
	if !tg() || tg() || !tg() {
		t.Errorf("toggler works incorrectly")
	}
}
