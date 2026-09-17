// if_x084: Проверка индекса
// Make the tests pass!
// I AM NOT DONE
//
// at должна вернуть элемент по индексу или "", если индекс вне диапазона.
// Для индекса, равного длине, функция паникует.
// Тренирует: границы допустимых индексов в условии.
// Сложность: hard
package main_test

import "testing"

func at(items []string, i int) string {
	if i >= 0 && i <= len(items) {
		return items[i]
	}
	return ""
}

func TestAt(t *testing.T) {
	items := []string{"a", "b", "c"}
	cases := map[int]string{0: "a", 2: "c", 3: "", -1: "", 10: ""}
	for in, want := range cases {
		if got := at(items, in); got != want {
			t.Errorf("at(%d) = %q, want %q", in, got, want)
		}
	}
}
