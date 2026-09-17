// range19
// Make the tests pass!

// I AM NOT DONE
//
// smooth заменяет каждый элемент (кроме первого) средним его и предыдущего
// элемента ИСХОДНОГО среза. Изменения на месте используют уже изменённые значения.
// Тренирует: range по срезу видит изменения последующих и предыдущих элементов.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func smooth(s []int) {
	for i := range s {
		if i > 0 {
			s[i] = (s[i] + s[i-1]) / 2
		}
	}
}

func TestSmooth(t *testing.T) {
	s := []int{0, 10, 20, 30}
	smooth(s)
	if !reflect.DeepEqual(s, []int{0, 5, 15, 25}) {
		t.Errorf("smooth = %v, want [0 5 15 25]", s)
	}
}
