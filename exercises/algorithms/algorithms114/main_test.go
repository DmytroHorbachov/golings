// algorithms114
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: быстрый и медленный указатели. Верните средний узел списка;
// при чётной длине — второй из двух средних.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func build(vals ...int) *ListNode {
	var head *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		head = &ListNode{vals[i], head}
	}
	return head
}

func middleNode(head *ListNode) *ListNode {
	return head
}

func TestMiddleNode(t *testing.T) {
	cases := []struct {
		vals []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 4},
		{[]int{9}, 9},
		{[]int{1, 2}, 2},
	}
	for _, c := range cases {
		if got := middleNode(build(c.vals...)); got == nil || got.Val != c.want {
			t.Errorf("middleNode(%v) = %v, want %d", c.vals, got, c.want)
		}
	}
	if middleNode(nil) != nil {
		t.Errorf("middleNode(nil) should be nil")
	}
}
