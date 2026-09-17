// algorithms40
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: рекурсия с границами. Проверьте, что дерево — корректное двоичное
// дерево поиска: слева строго меньше, справа строго больше.
// Сложность: medium. Ожидаемая асимптотика: O(n) по времени, O(h) по памяти
package main_test

import (
	"math"
	"testing"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return false
}

func TestIsValidBST(t *testing.T) {
	_ = math.MaxInt64
	valid := &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}
	if !isValidBST(valid) || !isValidBST(nil) {
		t.Errorf("valid BST rejected")
	}
	invalid := &TreeNode{5, &TreeNode{Val: 1}, &TreeNode{4, &TreeNode{Val: 3}, &TreeNode{Val: 6}}}
	if isValidBST(invalid) {
		t.Errorf("invalid BST accepted")
	}
	dup := &TreeNode{2, &TreeNode{Val: 2}, nil}
	if isValidBST(dup) {
		t.Errorf("duplicate values must be rejected")
	}
}
