// concurrent22
// Make the tests pass!

// I AM NOT DONE
//
// mergeCount читает из двух каналов по одному значению из каждого.
// Тренирует: select в цикле.
// Сложность: easy
package main_test

import "testing"

func mergeCount(a, b <-chan int) int {
	sum := 0
	for i := 0; i < 1; i++ {
		select {
		case v := <-a:
			sum += v
		case v := <-b:
			sum += v
		}
	}
	return sum
}

func TestMergeCount(t *testing.T) {
	a, b := make(chan int, 1), make(chan int, 1)
	a <- 3
	b <- 4
	if got := mergeCount(a, b); got != 7 {
		t.Errorf("mergeCount = %d", got)
	}
}
