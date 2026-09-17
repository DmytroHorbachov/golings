// arrays3
// Make the tests pass!

// I AM NOT DONE
//
// sparse должна вернуть массив длины 5, где только третий элемент (индекс 2) равен 10.
// Тренирует: литерал массива с явными индексами.
// Сложность: easy
package main_test

import "testing"

func sparse() [5]int {
	return [5]int{3: 10}
}

func TestSparse(t *testing.T) {
	if got := sparse(); got != [5]int{0, 0, 10, 0, 0} {
		t.Errorf("sparse = %v", got)
	}
}
