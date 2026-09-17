// algorithms82
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: быстрый и медленный указатели (Флойд). Определите, есть ли цикл
// в списке, не используя дополнительную память.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	return false
}

func TestHasCycle(t *testing.T) {
	nodes := make([]*ListNode, 4)
	for i := range nodes {
		nodes[i] = &ListNode{Val: i}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	if hasCycle(nodes[0]) {
		t.Errorf("list without cycle reported as cyclic")
	}
	nodes[3].Next = nodes[1]
	if !hasCycle(nodes[0]) {
		t.Errorf("cycle not detected")
	}
	self := &ListNode{}
	self.Next = self
	if !hasCycle(self) || hasCycle(nil) {
		t.Errorf("edge cases failed")
	}
}
