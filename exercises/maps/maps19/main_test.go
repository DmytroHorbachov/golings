// maps19
// Make the tests pass!

// I AM NOT DONE
//
// cachedSquare сначала ищет результат в кэше и только потом вычисляет.
// Тренирует: comma-ok при работе с кэшем.
// Сложность: easy
package main_test

import "testing"

var computed int

func cachedSquare(cache map[int]int, n int) int {
	if v, ok := cache[n]; ok {
		_ = v
	}
	computed++
	cache[n] = n * n
	return cache[n]
}

func TestCachedSquare(t *testing.T) {
	computed = 0
	cache := map[int]int{}
	if cachedSquare(cache, 4) != 16 || cachedSquare(cache, 4) != 16 {
		t.Errorf("wrong square")
	}
	if computed != 1 {
		t.Errorf("computed %d times, want 1", computed)
	}
}
