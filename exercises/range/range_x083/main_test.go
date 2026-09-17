// range_x083: Индекс после цикла
// Make the tests pass!
// I AM NOT DONE
//
// indexAfter возвращает позицию первого отрицательного числа или len(s).
// Код не компилируется: переменная цикла недоступна после него.
// Тренирует: переменные, объявленные в заголовке for, видны только внутри.
// Сложность: hard
package main_test

import "testing"

func indexAfter(s []int) int {
	for i, v := range s {
		if v < 0 {
			break
		}
	}
	return i
}

func TestIndexAfter(t *testing.T) {
	if indexAfter([]int{1, 2, -3, 4}) != 2 || indexAfter([]int{1, 2}) != 2 || indexAfter(nil) != 0 {
		t.Errorf("indexAfter works incorrectly")
	}
}
