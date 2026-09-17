// generics_x083: Рекурсивный обобщённый тип
// Make the tests pass!
// I AM NOT DONE
//
// Node[T] ссылается на себя, но без аргумента типа. Код не компилируется.
// Тренирует: внутри обобщённого типа ссылка на него требует [T].
// Сложность: hard
package main_test

import "testing"

type Node[T any] struct {
	Val  T
	Next *Node
}

func Len[T any](n *Node[T]) int {
	c := 0
	for ; n != nil; n = n.Next {
		c++
	}
	return c
}

func TestLen(t *testing.T) {
	list := &Node[string]{"a", &Node[string]{"b", nil}}
	if Len(list) != 2 {
		t.Errorf("Len = %d", Len(list))
	}
}
