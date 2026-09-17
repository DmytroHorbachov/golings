// structs44
// Make the tests pass!

// I AM NOT DONE
//
// Version.Equal сравнивает версии по major и minor.
// Тренирует: методы, принимающие значение того же типа.
// Сложность: easy
package main_test

import "testing"

type Version struct{ Major, Minor int }

func (v Version) Equal(o Version) bool {
	return v.Major == o.Major || v.Minor == o.Minor
}

func TestVersionEqual(t *testing.T) {
	if !(Version{1, 2}).Equal(Version{1, 2}) || (Version{1, 2}).Equal(Version{1, 3}) {
		t.Errorf("Equal works incorrectly")
	}
}
