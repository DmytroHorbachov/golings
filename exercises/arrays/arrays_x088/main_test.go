// arrays_x088: Порядок в ключе-массиве
// Make the tests pass!
// I AM NOT DONE
//
// friends хранит дружбу как map с ключом [2]string. Дружба симметрична,
// но поиск в обратном порядке имён не находит пару.
// Тренирует: массивы как ключи сравниваются поэлементно с учётом порядка.
// Сложность: hard
package main_test

import "testing"

type Friends map[[2]string]bool

func key(a, b string) [2]string {
	return [2]string{a, b}
}

func (f Friends) Add(a, b string)      { f[key(a, b)] = true }
func (f Friends) Are(a, b string) bool { return f[key(a, b)] }

func TestFriends(t *testing.T) {
	f := Friends{}
	f.Add("bob", "ann")
	if !f.Are("ann", "bob") || !f.Are("bob", "ann") {
		t.Errorf("friendship should be symmetric")
	}
	if f.Are("ann", "eve") {
		t.Errorf("ann and eve are not friends")
	}
}
