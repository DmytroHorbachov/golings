// anonymous_functions76
// Make the tests pass!

// I AM NOT DONE
//
// snapshotSum запоминает сумму набора в момент вызова, но литерал читает срез
// позже, когда данные уже изменились.
// Тренирует: литерал видит текущее содержимое среза, а не снимок.
// Сложность: hard
package main_test

import "testing"

func snapshotSum(data []int) func() int {
	return func() int {
		s := 0
		for _, v := range data {
			s += v
		}
		return s
	}
}

func TestSnapshotSum(t *testing.T) {
	data := []int{1, 2, 3}
	sum := snapshotSum(data)
	data[0] = 100
	if got := sum(); got != 6 {
		t.Errorf("sum = %d, want 6", got)
	}
}
