// slices_x075: Удаление при прямом обходе
// Make the tests pass!
// I AM NOT DONE
//
// dropNegatives удаляет отрицательные числа по индексу прямо во время обхода.
// Два отрицательных числа подряд удаляются не полностью.
// Тренирует: после удаления на место i встаёт следующий элемент.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func dropNegatives(s []int) []int {
	for i := 0; i < len(s); i++ {
		if s[i] < 0 {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func TestDropNegatives(t *testing.T) {
	if got := dropNegatives([]int{1, -2, -3, 4, -5}); !reflect.DeepEqual(got, []int{1, 4}) {
		t.Errorf("dropNegatives = %v", got)
	}
}
