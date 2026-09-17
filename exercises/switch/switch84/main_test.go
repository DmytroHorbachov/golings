// switch84
// Make the tests pass!

// I AM NOT DONE
//
// render выводит Money особым образом, а остальные fmt.Stringer — через String().
// Money тоже реализует Stringer, и до своей ветки не доходит.
// Тренирует: ветки type switch проверяются по порядку, интерфейсы «перехватывают» типы.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

type Money struct{ Cents int }

func (m Money) String() string { return fmt.Sprintf("%d cents", m.Cents) }

type Tag string

func (t Tag) String() string { return "#" + string(t) }

func render(v interface{}) string {
	switch x := v.(type) {
	case fmt.Stringer:
		return x.String()
	case Money:
		return fmt.Sprintf("$%d.%02d", x.Cents/100, x.Cents%100)
	}
	return "?"
}

func TestRender(t *testing.T) {
	if got := render(Money{1234}); got != "$12.34" {
		t.Errorf("render(Money) = %q, want $12.34", got)
	}
	if got := render(Tag("go")); got != "#go" {
		t.Errorf("render(Tag) = %q, want #go", got)
	}
}
