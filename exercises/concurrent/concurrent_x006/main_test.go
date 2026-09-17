// concurrent_x006: Получение результата
// Make the tests pass!
// I AM NOT DONE
//
// compute запускает вычисление в горутине, но не читает результат из канала.
// Тренирует: получение значения из канала.
// Сложность: easy
package main_test

import "testing"

func compute(a, b int) int {
	ch := make(chan int, 1)
	go func() { ch <- a * b }()
	return len(ch)
}

func TestCompute(t *testing.T) {
	if got := compute(6, 7); got != 42 {
		t.Errorf("compute = %d", got)
	}
}
