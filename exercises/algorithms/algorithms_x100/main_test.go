// algorithms_x100: Number of Islands II (острова по мере добавления)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: динамический union-find. Клетки добавляются по одной; после каждой
// операции верните текущее количество островов.
// Сложность: hard. Ожидаемая асимптотика: O(k·α(rc)) по времени, O(r·c) по памяти
package main_test

import (
	"reflect"
	"testing"
)

func numIslands2(rows, cols int, positions [][2]int) []int {
	return nil
}

func TestNumIslands2(t *testing.T) {
	got := numIslands2(3, 3, [][2]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}})
	if !reflect.DeepEqual(got, []int{1, 1, 2, 3}) {
		t.Errorf("numIslands2 = %v, want [1 1 2 3]", got)
	}
	got = numIslands2(1, 1, [][2]int{{0, 0}, {0, 0}})
	if !reflect.DeepEqual(got, []int{1, 1}) {
		t.Errorf("duplicate position = %v", got)
	}
	if got := numIslands2(2, 2, nil); len(got) != 0 {
		t.Errorf("no positions = %v", got)
	}
}
