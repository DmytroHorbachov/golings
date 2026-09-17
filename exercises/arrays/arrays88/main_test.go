// arrays88
// Make the tests pass!

// I AM NOT DONE
//
// ageBuckets раскладывает возрасты по корзинам по 10 лет: 0–9, 10–19, ..., 90+.
// Тренирует: вычисление индекса и ограничение последней корзиной.
// Сложность: medium
package main_test

import "testing"

func ageBuckets(ages []int) [10]int {
	var b [10]int
	for _, a := range ages {
		b[a%10]++
	}
	return b
}

func TestAgeBuckets(t *testing.T) {
	got := ageBuckets([]int{5, 15, 19, 42, 99, 120, -1})
	want := [10]int{1, 2, 0, 0, 1, 0, 0, 0, 0, 2}
	if got != want {
		t.Errorf("ageBuckets = %v, want %v", got, want)
	}
}
