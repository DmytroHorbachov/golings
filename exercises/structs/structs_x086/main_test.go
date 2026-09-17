// structs_x086: Рекурсивный тип
// Make the tests pass!
// I AM NOT DONE
//
// Category содержит родительскую категорию. Код не компилируется:
// структура не может содержать саму себя по значению.
// Тренирует: рекурсивные структуры требуют указателей.
// Сложность: hard
package main_test

import "testing"

type Category struct {
	Name   string
	Parent Category
}

func path(c *Category) string {
	if c.Parent == nil {
		return c.Name
	}
	return path(c.Parent) + "/" + c.Name
}

func TestPath(t *testing.T) {
	root := &Category{Name: "shop"}
	kids := &Category{Name: "toys", Parent: root}
	if got := path(kids); got != "shop/toys" {
		t.Errorf("path = %q", got)
	}
}
