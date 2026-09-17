// anonymous_functions101
// Make the tests pass!

// I AM NOT DONE
//
// compute запускает вычисление в горутине и получает результат из канала.
// Тренирует: go func() { ... }().
// Сложность: easy
package main_test

import "testing"

func compute(a, b int) int {
	ch := make(chan int)
	go func() {
		ch <- a
	}()
	return <-ch
}

func TestCompute(t *testing.T) {
	if compute(6, 7) != 42 {
		t.Errorf("compute = %d", compute(6, 7))
	}
}
