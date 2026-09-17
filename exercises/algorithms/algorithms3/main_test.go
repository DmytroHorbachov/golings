// algorithms3
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: середина списка и разворот половины. Проверьте, читается ли список
// одинаково в обе стороны.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) доп. памяти
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

func isPalindromeList(head *ListNode) bool {
	return false
}

func TestIsPalindromeList(t *testing.T) {
	cases := []struct {
		vals []int
		want bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{7}, true},
		{nil, true},
		{[]int{1, 2, 3, 1}, false},
	}
	for _, c := range cases {
		if got := isPalindromeList(build(c.vals...)); got != c.want {
			t.Errorf("isPalindromeList(%v) = %v, want %v", c.vals, got, c.want)
		}
	}
}
