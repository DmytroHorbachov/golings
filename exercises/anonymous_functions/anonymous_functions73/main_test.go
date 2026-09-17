// anonymous_functions73
// Make the tests pass!

// I AM NOT DONE
//
// makeGetters создаёт функции, каждая из которых должна вернуть свой номер.
// Все функции возвращают одно и то же число.
// Тренирует: литералы захватывают переменную, объявленную вне цикла, а не её значение.
// Сложность: hard
package main_test

import "testing"

func makeGetters(n int) []func() int {
	var fs []func() int
	i := 0
	for i < n {
		fs = append(fs, func() int { return i })
		i++
	}
	return fs
}

func TestMakeGetters(t *testing.T) {
	fs := makeGetters(3)
	for want, f := range fs {
		if got := f(); got != want {
			t.Errorf("getter %d returned %d", want, got)
		}
	}
}
