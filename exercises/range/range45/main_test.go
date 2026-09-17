// range45
// Make the tests pass!

// I AM NOT DONE
//
// sumOrZero суммирует элементы массива по указателю, а для nil возвращает 0.
// Тренирует: range с переменной значения по nil *[N]T паникует.
// Сложность: hard
package main_test

import "testing"

func sumOrZero(p *[4]int) int {
	s := 0
	for _, v := range p {
		s += v
	}
	return s
}

func TestSumOrZero(t *testing.T) {
	if got := sumOrZero(&[4]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("sumOrZero = %d", got)
	}
	if got := sumOrZero(nil); got != 0 {
		t.Errorf("sumOrZero(nil) = %d", got)
	}
}
