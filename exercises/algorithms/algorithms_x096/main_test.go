// algorithms_x096: Longest Increasing Path in a Matrix (самый длинный возрастающий путь)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: поиск в глубину с мемоизацией. Найдите длину самого длинного пути
// по соседним клеткам, где значения строго возрастают.
// Сложность: hard. Ожидаемая асимптотика: O(r·c) по времени, O(r·c) по памяти
package main_test

import "testing"

func longestIncreasingPath(m [][]int) int {
	return 0
}

func TestLongestIncreasingPath(t *testing.T) {
	cases := []struct {
		m    [][]int
		want int
	}{
		{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
		{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, 4},
		{[][]int{{1}}, 1},
		{nil, 0},
		{[][]int{{7, 7}, {7, 7}}, 1},
	}
	for _, c := range cases {
		if got := longestIncreasingPath(c.m); got != c.want {
			t.Errorf("longestIncreasingPath(%v) = %d, want %d", c.m, got, c.want)
		}
	}
}
