// range_x013: break на стоп-значении
// Make the tests pass!
// I AM NOT DONE
//
// readUntilZero складывает числа до первого нуля (не включая его и всё после).
// Тренирует: break внутри range.
// Сложность: easy
package main_test

import "testing"

func readUntilZero(s []int) int {
	total := 0
	for _, v := range s {
		if v == 0 {
			continue
		}
		total += v
	}
	return total
}

func TestReadUntilZero(t *testing.T) {
	if got := readUntilZero([]int{4, 5, 0, 100}); got != 9 {
		t.Errorf("readUntilZero = %d, want 9", got)
	}
}
