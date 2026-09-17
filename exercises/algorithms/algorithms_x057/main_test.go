// algorithms_x057: Invert Binary Tree (зеркальное дерево)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: рекурсия по дереву. Поменяйте местами левое и правое поддеревья
// на всех уровнях и верните корень.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(h) по памяти
package main_test

import (
	"reflect"
	"testing"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inorder(t *TreeNode, out []int) []int {
	if t == nil {
		return out
	}
	out = inorder(t.Left, out)
	out = append(out, t.Val)
	return inorder(t.Right, out)
}

func invertTree(root *TreeNode) *TreeNode {
	return root
}

func TestInvertTree(t *testing.T) {
	tree := &TreeNode{4,
		&TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
		&TreeNode{7, &TreeNode{Val: 6}, &TreeNode{Val: 9}}}
	got := inorder(invertTree(tree), nil)
	if !reflect.DeepEqual(got, []int{9, 7, 6, 4, 3, 2, 1}) {
		t.Errorf("inverted inorder = %v", got)
	}
	if invertTree(nil) != nil {
		t.Errorf("invertTree(nil) should be nil")
	}
}
