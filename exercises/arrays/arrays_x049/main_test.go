// arrays_x049: Удаление дубликатов
// Make the tests pass!
// I AM NOT DONE
//
// dedupe убирает повторы из отсортированного массива на месте и возвращает
// количество уникальных элементов.
// Тренирует: запись в массив по отдельному индексу.
// Сложность: medium
package main_test

import "testing"

func dedupe(a *[7]int) int {
	w := 1
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			w++
		}
	}
	return w
}

func TestDedupe(t *testing.T) {
	a := [7]int{1, 1, 2, 3, 3, 3, 4}
	n := dedupe(&a)
	if n != 4 || a[0] != 1 || a[1] != 2 || a[2] != 3 || a[3] != 4 {
		t.Errorf("dedupe = %d, %v", n, a)
	}
}
