// structs_x039: Связный список
// Make the tests pass!
// I AM NOT DONE
//
// List.PushFront добавляет значение в начало, Values возвращает все значения.
// Тренирует: структуры со ссылками на себя.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type node struct {
	val  int
	next *node
}

type List struct{ head *node }

func (l *List) PushFront(v int) {
	l.head = &node{val: v}
}

func (l *List) Values() []int {
	var out []int
	for n := l.head; n != nil; n = nil {
		out = append(out, n.val)
	}
	return out
}

func TestList(t *testing.T) {
	var l List
	l.PushFront(3)
	l.PushFront(2)
	l.PushFront(1)
	if !reflect.DeepEqual(l.Values(), []int{1, 2, 3}) {
		t.Errorf("Values = %v", l.Values())
	}
}
