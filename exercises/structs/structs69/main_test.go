// structs69
// Make the tests pass!

// I AM NOT DONE
//
// Report встраивает Author и Editor, у обоих есть поле Name.
// Код не компилируется: r.Name неоднозначно.
// Тренирует: конфликт продвинутых полей на одной глубине.
// Сложность: hard
package main_test

import "testing"

type Author struct{ Name string }
type Editor struct{ Name string }

type Report struct {
	Author
	Editor
}

func byline(r Report) string {
	return "by " + r.Name
}

func TestByline(t *testing.T) {
	r := Report{Author{"ann"}, Editor{"bob"}}
	if byline(r) != "by ann" {
		t.Errorf("byline = %q", byline(r))
	}
}
