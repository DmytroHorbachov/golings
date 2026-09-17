// arrays_x085: Массив в вариативную функцию
// Make the tests pass!
// I AM NOT DONE
//
// total передаёт элементы массива в вариативную функцию sum.
// Код не компилируется.
// Тренирует: распаковывать через ... можно только срез.
// Сложность: hard
package main_test

import "testing"

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func total(a [4]int) int {
	return sum(a...)
}

func TestTotal(t *testing.T) {
	if got := total([4]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("total = %d", got)
	}
}
