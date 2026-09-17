// slices52
// Make the tests pass!

// I AM NOT DONE
//
// expand проходит по очереди задач и добавляет подзадачи в её конец;
// подзадачи тоже должны быть обработаны. Сейчас они пропускаются.
// Тренирует: range вычисляет длину среза один раз перед началом цикла.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

var children = map[string][]string{
	"root": {"a", "b"},
	"a":    {"a1"},
}

func expand(start string) []string {
	queue := []string{start}
	var visited []string
	for _, n := range queue {
		visited = append(visited, n)
		queue = append(queue, children[n]...)
	}
	return visited
}

func TestExpand(t *testing.T) {
	if got := expand("root"); !reflect.DeepEqual(got, []string{"root", "a", "b", "a1"}) {
		t.Errorf("expand = %v", got)
	}
}
