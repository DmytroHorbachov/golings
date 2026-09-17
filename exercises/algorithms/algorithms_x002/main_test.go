// algorithms_x002: Contains Duplicate (есть ли повтор)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: хэш-таблицы. Верните true, если хотя бы одно значение
// встречается в срезе больше одного раза.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import "testing"

func containsDuplicate(nums []int) bool {
	return false
}

func TestContainsDuplicate(t *testing.T) {
	big := make([]int, 100000)
	for i := range big {
		big[i] = i
	}
	cases := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{nil, false},
		{[]int{7}, false},
		{[]int{-5, 5, -5}, true},
		{big, false},
		{append(big, 99999), true},
	}
	for _, c := range cases {
		if got := containsDuplicate(c.nums); got != c.want {
			t.Errorf("containsDuplicate(len=%d) = %v, want %v", len(c.nums), got, c.want)
		}
	}
}
