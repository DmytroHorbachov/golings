// structs_x034: Пустая структура как значение
// Make the tests pass!
// I AM NOT DONE
//
// Множество тегов реализовано через map[string]struct{}.
// Тренирует: пустую структуру struct{}.
// Сложность: easy
package main_test

import "testing"

type TagSet map[string]struct{}

func (s TagSet) Add(tag string) {
	s[""] = struct{}{}
}

func TestTagSet(t *testing.T) {
	s := TagSet{}
	s.Add("go")
	s.Add("go")
	if _, ok := s["go"]; !ok || len(s) != 1 {
		t.Errorf("set = %v", s)
	}
}
