// arrays47
// Make the tests pass!

// I AM NOT DONE
//
// prev возвращает индекс предыдущей позиции в кольцевом массиве из 5 элементов.
// Для позиции 0 получается -1 и программа паникует.
// Тренирует: остаток от деления отрицательного числа при индексации массива.
// Сложность: hard
package main_test

import "testing"

var ring = [5]string{"a", "b", "c", "d", "e"}

func prev(i int) string {
	return ring[(i-1)%len(ring)]
}

func TestPrev(t *testing.T) {
	if prev(3) != "c" || prev(0) != "e" || prev(1) != "a" {
		t.Errorf("prev works incorrectly")
	}
}
