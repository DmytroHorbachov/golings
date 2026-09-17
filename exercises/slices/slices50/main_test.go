// slices50
// Make the tests pass!

// I AM NOT DONE
//
// Pool отдаёт буферы для повторного использования. grow расширяет буфер
// до нужной длины, но в новых элементах остаются чужие старые данные.
// Тренирует: s[:n] при n <= cap открывает прежнее содержимое массива.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func grow(s []int, n int) []int {
	old := len(s)
	if n <= cap(s) {
		s = s[:n]
		return s
	}
	out := make([]int, n)
	copy(out, s)
	return out
}

func TestGrow(t *testing.T) {
	buf := []int{7, 7, 7, 7}
	buf = buf[:1]
	got := grow(buf, 3)
	if !reflect.DeepEqual(got, []int{7, 0, 0}) {
		t.Errorf("grow = %v, want [7 0 0]", got)
	}
}
