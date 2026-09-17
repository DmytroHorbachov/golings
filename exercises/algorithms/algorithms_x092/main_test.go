// algorithms_x092: Find if Path Exists (существует ли путь)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: обход в глубину или ширину. Определите, есть ли путь между
// вершинами start и end в неориентированном графе.
// Сложность: easy. Ожидаемая асимптотика: O(V + E) по времени, O(V + E) по памяти
package main_test

import "testing"

func validPath(n int, edges [][2]int, start, end int) bool {
	return false
}

func TestValidPath(t *testing.T) {
	if !validPath(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2) {
		t.Errorf("path should exist")
	}
	if validPath(6, [][2]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5) {
		t.Errorf("path should not exist")
	}
	if !validPath(1, nil, 0, 0) {
		t.Errorf("start == end")
	}
}
