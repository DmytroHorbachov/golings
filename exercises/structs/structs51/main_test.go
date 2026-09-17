// structs51
// Make the tests pass!

// I AM NOT DONE
//
// Clone должен вернуть полностью независимую копию плейлиста.
// Тренирует: копирование структуры с полем-срезом.
// Сложность: medium
package main_test

import "testing"

type Playlist struct {
	Name  string
	Songs []string
}

func (p Playlist) Clone() Playlist {
	return p
}

func TestClone(t *testing.T) {
	p := Playlist{"rock", []string{"a", "b"}}
	c := p.Clone()
	c.Songs[0] = "z"
	c.Name = "pop"
	if p.Songs[0] != "a" || p.Name != "rock" {
		t.Errorf("original changed: %+v", p)
	}
}
