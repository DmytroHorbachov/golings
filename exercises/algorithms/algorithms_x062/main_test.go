// algorithms_x062: Kth Smallest Element in a BST (k-е по возрастанию)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: симметричный обход BST. Верните k-е наименьшее значение (k с единицы).
// Если узлов меньше k — верните -1.
// Сложность: medium. Ожидаемая асимптотика: O(h + k) по времени, O(h) по памяти
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	return 0
}

func TestKthSmallest(t *testing.T) {
	tree := &TreeNode{5,
		&TreeNode{3, &TreeNode{Val: 2}, &TreeNode{Val: 4}},
		&TreeNode{Val: 6}}
	cases := map[int]int{1: 2, 3: 4, 5: 6, 6: -1}
	for k, want := range cases {
		if got := kthSmallest(tree, k); got != want {
			t.Errorf("kthSmallest(%d) = %d, want %d", k, got, want)
		}
	}
	if kthSmallest(nil, 1) != -1 {
		t.Errorf("empty tree should give -1")
	}
}
