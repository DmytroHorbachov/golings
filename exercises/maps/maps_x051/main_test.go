// maps_x051: Степени вершин
// Make the tests pass!
// I AM NOT DONE
//
// degrees считает степень каждой вершины неориентированного графа по списку рёбер.
// Тренирует: увеличение счётчиков для обоих концов ребра.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func degrees(edges [][2]string) map[string]int {
	d := map[string]int{}
	for _, e := range edges {
		d[e[0]] += 2
	}
	return d
}

func TestDegrees(t *testing.T) {
	got := degrees([][2]string{{"a", "b"}, {"b", "c"}, {"c", "c"}})
	if !reflect.DeepEqual(got, map[string]int{"a": 1, "b": 2, "c": 3}) {
		t.Errorf("degrees = %v", got)
	}
}
