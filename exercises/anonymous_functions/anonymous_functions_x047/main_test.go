// anonymous_functions_x047: Посетитель с пропуском поддеревьев
// Make the tests pass!
// I AM NOT DONE
//
// walk обходит дерево; литерал-посетитель возвращает false, чтобы не заходить
// в детей узла. collect не должен заходить в ветки с именем "skip".
// Тренирует: литерал, управляющий обходом.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type Node struct {
	Name     string
	Children []*Node
}

func walk(n *Node, visit func(*Node) bool) {
	visit(n)
	for _, c := range n.Children {
		walk(c, visit)
	}
}

func collect(root *Node) []string {
	var names []string
	walk(root, func(n *Node) bool {
		if n.Name == "skip" {
			return false
		}
		names = append(names, n.Name)
		return true
	})
	return names
}

func TestCollect(t *testing.T) {
	root := &Node{"root", []*Node{
		{"a", nil},
		{"skip", []*Node{{"hidden", nil}}},
		{"b", []*Node{{"c", nil}}},
	}}
	if got := collect(root); !reflect.DeepEqual(got, []string{"root", "a", "b", "c"}) {
		t.Errorf("collect = %v", got)
	}
}
