// range_x100: Ключи map как индексы
// Make the tests pass!
// I AM NOT DONE
//
// toSlice превращает map «номер страницы -> заголовок» в срез заголовков,
// упорядоченный по номеру страницы. Номера идут с пропусками, и код паникует.
// Тренирует: ключи map — не позиции; порядок нужно получить сортировкой.
// Сложность: hard
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func toSlice(pages map[int]string) []string {
	out := make([]string, len(pages))
	for n, title := range pages {
		out[n] = title
	}
	return out
}

func TestToSlice(t *testing.T) {
	_ = sort.Ints
	got := toSlice(map[int]string{10: "c", 1: "a", 5: "b"})
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("toSlice = %v", got)
	}
}
