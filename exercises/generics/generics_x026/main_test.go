// generics_x026: Значение по умолчанию
// Make the tests pass!
// I AM NOT DONE
//
// OrDefault возвращает def, если v равно нулевому значению типа.
// Тренирует: сравнение с нулевым значением через comparable.
// Сложность: easy
package main_test

import "testing"

func OrDefault[T comparable](v, def T) T {
	var zero T
	if v != zero {
		return def
	}
	return v
}

func TestOrDefault(t *testing.T) {
	if OrDefault("", "guest") != "guest" || OrDefault("ann", "guest") != "ann" || OrDefault(0, 8080) != 8080 {
		t.Errorf("OrDefault works incorrectly")
	}
}
