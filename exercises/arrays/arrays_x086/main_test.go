// arrays_x086: append к срезу массива
// Make the tests pass!
// I AM NOT DONE
//
// extend добавляет элемент к срезу полного массива и затем меняет первый элемент.
// Ожидается, что изменения попадут в исходный массив, но этого не происходит.
// Тренирует: append при нехватке ёмкости создаёт новый массив.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func extend(a *[3]int, v int) []int {
	s := a[:]
	s = append(s, v)
	s[0] = -1
	return s
}

func TestExtend(t *testing.T) {
	a := [3]int{1, 2, 3}
	s := extend(&a, 4)
	if a[0] != -1 {
		t.Errorf("array not updated: %v", a)
	}
	if !reflect.DeepEqual(s, []int{-1, 2, 3, 4}) {
		t.Errorf("slice = %v", s)
	}
}
