// slices_x099: Сдвиг вправо
// Make the tests pass!
// I AM NOT DONE
//
// shiftRight сдвигает элементы на одну позицию вправо (последний теряется),
// а первый элемент становится нулём. Сейчас весь срез заполняется первым значением.
// Тренирует: порядок копирования при перекрывающихся областях.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func shiftRight(s []int) {
	for i := 1; i < len(s); i++ {
		s[i] = s[i-1]
	}
	if len(s) > 0 {
		s[0] = 0
	}
}

func TestShiftRight(t *testing.T) {
	s := []int{1, 2, 3, 4}
	shiftRight(s)
	if !reflect.DeepEqual(s, []int{0, 1, 2, 3}) {
		t.Errorf("shiftRight = %v", s)
	}
}
