// structs_x096: Значимый метод у nil-указателя
// Make the tests pass!
// I AM NOT DONE
//
// Метод Label со значимым получателем вызывается у nil-указателя и паникует
// (в отличие от методов с указателем-получателем).
// Тренирует: вызов value-метода через указатель разыменовывает его.
// Сложность: hard
package main_test

import "testing"

type Tag struct{ Name string }

func (t Tag) Label() string {
	return "#" + t.Name
}

func TestLabel(t *testing.T) {
	var none *Tag
	if none.Label() != "#untagged" || (&Tag{"go"}).Label() != "#go" {
		t.Errorf("Label works incorrectly")
	}
}
