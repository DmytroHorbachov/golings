// structs66
// Make the tests pass!

// I AM NOT DONE
//
// title возвращает заголовок или "untitled" для nil.
// Тренирует: проверку указателя на структуру.
// Сложность: easy
package main_test

import "testing"

type Post struct{ Title string }

func title(p *Post) string {
	if p != nil {
		return "untitled"
	}
	return p.Title
}

func TestTitle(t *testing.T) {
	if title(nil) != "untitled" || title(&Post{"Go"}) != "Go" {
		t.Errorf("title works incorrectly")
	}
}
