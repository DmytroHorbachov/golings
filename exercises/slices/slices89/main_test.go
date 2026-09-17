// slices89
// Make the tests pass!

// I AM NOT DONE
//
// Тип Stack основан на срезе. Метод Push должен добавлять элемент,
// но стек остаётся пустым.
// Тренирует: append в методе со значимым получателем меняет копию заголовка.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

type Stack []int

func (s Stack) Push(v int) {
	s = append(s, v)
}

func TestStackPush(t *testing.T) {
	var s Stack
	s.Push(1)
	s.Push(2)
	if !reflect.DeepEqual([]int(s), []int{1, 2}) {
		t.Errorf("stack = %v, want [1 2]", s)
	}
}
