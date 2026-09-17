// arrays59
// Make the tests pass!

// I AM NOT DONE
//
// defaults должна вернуть массив [1 2 0 0]: незаданные элементы равны нулю.
// Тренирует: частичную инициализацию массива.
// Сложность: easy
package main_test

import "testing"

func defaults() [4]int {
	return [4]int{1, 2, 2, 2}
}

func TestDefaults(t *testing.T) {
	if got := defaults(); got != [4]int{1, 2, 0, 0} {
		t.Errorf("defaults = %v", got)
	}
}
