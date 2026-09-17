// arrays10
// Make the tests pass!

// I AM NOT DONE
//
// Названия уровней хранятся в массиве, индексом служит константа уровня.
// Тренирует: массив с индексированным литералом по константам.
// Сложность: easy
package main_test

import "testing"

const (
	Low = iota
	Mid
	High
)

var levelNames = [...]string{
	Low:  "low",
	Mid:  "high",
	High: "high",
}

func TestLevelNames(t *testing.T) {
	if levelNames[Low] != "low" || levelNames[Mid] != "mid" || levelNames[High] != "high" {
		t.Errorf("levelNames = %v", levelNames)
	}
}
