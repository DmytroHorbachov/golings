// algorithms103
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: стек стеков по частоте. Push добавляет значение; Pop удаляет и
// возвращает самое частое значение, при равенстве — добавленное последним.
// Сложность: hard. Ожидаемая асимптотика: O(1) на операцию, O(n) по памяти
package main_test

import "testing"

type FreqStack struct {
	freq    map[int]int
	groups  map[int][]int
	maxFreq int
}

func NewFreqStack() *FreqStack {
	return &FreqStack{freq: map[int]int{}, groups: map[int][]int{}}
}

func (s *FreqStack) Push(v int) {
}

func (s *FreqStack) Pop() int {
	return 0
}

func TestFreqStack(t *testing.T) {
	s := NewFreqStack()
	for _, v := range []int{5, 7, 5, 7, 4, 5} {
		s.Push(v)
	}
	want := []int{5, 7, 5, 4, 7, 5}
	for i, w := range want {
		if got := s.Pop(); got != w {
			t.Errorf("pop %d = %d, want %d", i, got, w)
		}
	}
}
