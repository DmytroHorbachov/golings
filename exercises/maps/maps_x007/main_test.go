// maps_x007: Пустой литерал
// Make the tests pass!
// I AM NOT DONE
//
// emptyIndex возвращает пустую, но готовую к записи map.
// Тренирует: инициализацию map литералом {}.
// Сложность: easy
package main_test

import "testing"

func emptyIndex() map[string][]int {
	idx := map[string][]int{}
	idx["init"] = nil
	return idx
}

func TestEmptyIndex(t *testing.T) {
	idx := emptyIndex()
	if idx == nil || len(idx) != 0 {
		t.Fatalf("emptyIndex = %v", idx)
	}
	idx["go"] = []int{1}
}
