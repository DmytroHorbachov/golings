// slices_x093: append к значению из map
// Make the tests pass!
// I AM NOT DONE
//
// addTag добавляет тег в список тегов статьи, хранящийся в map.
// Добавленный тег не сохраняется.
// Тренирует: append возвращает новый заголовок среза, и его нужно записать обратно в map.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func addTag(m map[string][]string, post, tag string) {
	tags := m[post]
	tags = append(tags, tag)
}

func TestAddTag(t *testing.T) {
	m := map[string][]string{}
	addTag(m, "p1", "go")
	addTag(m, "p1", "slices")
	if !reflect.DeepEqual(m["p1"], []string{"go", "slices"}) {
		t.Errorf("m[p1] = %v", m["p1"])
	}
}
