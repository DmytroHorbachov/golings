// functions_x091: nil-получатель
// Make the tests pass!
// I AM NOT DONE
//
// Sum должна вернуть сумму значений связного списка; пустой список — это nil.
// Сейчас вызов на пустом списке паникует.
// Тренирует: методы с указателем-получателем можно вызывать на nil.
// Сложность: hard
package main_test

import "testing"

type List struct {
	Val  int
	Next *List
}

func (l *List) Sum() int {
	if l.Next == nil {
		return l.Val
	}
	return l.Val + l.Next.Sum()
}

func TestListSum(t *testing.T) {
	var empty *List
	if got := empty.Sum(); got != 0 {
		t.Errorf("empty.Sum() = %d, want 0", got)
	}
	l := &List{1, &List{2, &List{3, nil}}}
	if got := l.Sum(); got != 6 {
		t.Errorf("Sum() = %d, want 6", got)
	}
}
