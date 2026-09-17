// structs_x095: Метод встроенного типа меняет не то поле
// Make the tests pass!
// I AM NOT DONE
//
// Product встраивает Base (с полем Title) и имеет собственное поле Title.
// Метод Base.Rename, вызванный у Product, меняет Base.Title, а не Product.Title.
// Тренирует: продвинутый метод работает со встроенным значением, а не с внешним.
// Сложность: hard
package main_test

import "testing"

type Base struct{ Title string }

func (b *Base) Rename(t string) { b.Title = t }

type Product struct {
	Base
	Title string
}

func TestRename(t *testing.T) {
	p := &Product{Title: "old"}
	p.Rename("new")
	if p.Title != "new" {
		t.Errorf("Title = %q, want new", p.Title)
	}
}
