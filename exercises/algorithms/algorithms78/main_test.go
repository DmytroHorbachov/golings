// algorithms78
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: рекурсия с побочным результатом. Путь — последовательность узлов,
// соединённых рёбрами, идущая в любом направлении, но без повторов. Найдите
// максимальную сумму значений такого пути (дерево непустое, значения любые).
// Сложность: hard. Ожидаемая асимптотика: O(n) по времени, O(h) по памяти
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	return 0
}

func TestMaxPathSum(t *testing.T) {
	cases := []struct {
		tree *TreeNode
		want int
	}{
		{&TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}, 6},
		{&TreeNode{-10, &TreeNode{Val: 9}, &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}, 42},
		{&TreeNode{Val: -3}, -3},
		{&TreeNode{-2, &TreeNode{Val: -1}, nil}, -1},
	}
	for i, c := range cases {
		if got := maxPathSum(c.tree); got != c.want {
			t.Errorf("case %d: maxPathSum = %d, want %d", i, got, c.want)
		}
	}
}
