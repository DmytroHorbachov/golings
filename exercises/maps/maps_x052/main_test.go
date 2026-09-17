// maps_x052: Список смежности
// Make the tests pass!
// I AM NOT DONE
//
// adjacency строит список смежности ориентированного графа,
// соседи каждой вершины отсортированы.
// Тренирует: map[string][]string.
// Сложность: medium
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func adjacency(edges [][2]string) map[string][]string {
	adj := map[string][]string{}
	for _, e := range edges {
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	return adj
}

func TestAdjacency(t *testing.T) {
	got := adjacency([][2]string{{"a", "c"}, {"a", "b"}, {"b", "c"}})
	if !reflect.DeepEqual(got, map[string][]string{"a": {"b", "c"}, "b": {"c"}}) {
		t.Errorf("adjacency = %v", got)
	}
}
