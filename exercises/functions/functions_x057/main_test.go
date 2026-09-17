// functions_x057: Обход с посетителем
// Make the tests pass!
// I AM NOT DONE
//
// walk обходит дерево в прямом порядке и вызывает visit для каждого узла.
// Сейчас обход пропускает правые поддеревья и посещает узел после детей.
// Тренирует: callback-функции и рекурсию по дереву.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type Node struct {
	Val         int
	Left, Right *Node
}

func walk(n *Node, visit func(int)) {
	if n == nil {
		return
	}
	walk(n.Left, visit)
	visit(n.Val)
}

func TestWalk(t *testing.T) {
	root := &Node{1, &Node{2, &Node{4, nil, nil}, nil}, &Node{3, nil, &Node{5, nil, nil}}}
	var got []int
	walk(root, func(v int) { got = append(got, v) })
	if want := []int{1, 2, 4, 3, 5}; !reflect.DeepEqual(got, want) {
		t.Errorf("walk = %v, want %v", got, want)
	}
}
