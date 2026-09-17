// algorithms_x056: Maximum Depth of Binary Tree (глубина дерева)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: рекурсия по дереву. Верните количество узлов на самом длинном пути
// от корня до листа. У пустого дерева глубина 0.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(h) по памяти
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return 0
}

func TestMaxDepth(t *testing.T) {
	tree := &TreeNode{3,
		&TreeNode{Val: 9},
		&TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}
	if got := maxDepth(tree); got != 3 {
		t.Errorf("maxDepth = %d, want 3", got)
	}
	if maxDepth(nil) != 0 || maxDepth(&TreeNode{Val: 1}) != 1 {
		t.Errorf("edge cases failed")
	}
	deep := &TreeNode{Val: 0}
	cur := deep
	for i := 0; i < 999; i++ {
		cur.Left = &TreeNode{Val: i}
		cur = cur.Left
	}
	if got := maxDepth(deep); got != 1000 {
		t.Errorf("deep tree depth = %d, want 1000", got)
	}
}
