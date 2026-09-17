// maps_x042: Индекс по id
// Make the tests pass!
// I AM NOT DONE
//
// indexByID строит map от id к указателю на элемент исходного среза.
// Изменения через map должны быть видны в срезе.
// Тренирует: map[int]*T и адреса элементов среза.
// Сложность: medium
package main_test

import "testing"

type Item struct {
	ID    int
	Title string
}

func indexByID(items []Item) map[int]*Item {
	idx := make(map[int]*Item, len(items))
	for _, it := range items {
		copyIt := it
		idx[it.ID] = &copyIt
	}
	return idx
}

func TestIndexByID(t *testing.T) {
	items := []Item{{1, "a"}, {2, "b"}}
	idx := indexByID(items)
	idx[2].Title = "B"
	if items[1].Title != "B" {
		t.Errorf("items[1].Title = %q, want B", items[1].Title)
	}
}
