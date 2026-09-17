// algorithms_x045: Reverse Linked List (разворот списка)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: связные списки. Разверните односвязный список: последний узел
// должен стать головой. Верните новую голову списка.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import (
	"reflect"
	"testing"
)

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

func toSlice(l *ListNode) []int {
	out := []int{}
	for ; l != nil; l = l.Next {
		out = append(out, l.Val)
	}
	return out
}

func reverseList(head *ListNode) *ListNode {
	return head
}

func TestReverseList(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{7}, []int{7}},
		{nil, []int{}},
	}
	for _, c := range cases {
		if got := toSlice(reverseList(build(c.in...))); !reflect.DeepEqual(got, c.want) {
			t.Errorf("reverseList(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
