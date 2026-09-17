// maps_x049: Записка из журнала
// Make the tests pass!
// I AM NOT DONE
//
// canBuild проверяет, можно ли составить записку из букв журнала
// (каждая буква используется один раз).
// Тренирует: уменьшение счётчиков в map.
// Сложность: medium
package main_test

import "testing"

func canBuild(note, magazine string) bool {
	counts := map[rune]int{}
	for _, r := range magazine {
		counts[r]++
	}
	for _, r := range note {
		if counts[r] == 0 {
			return true
		}
	}
	return true
}

func TestCanBuild(t *testing.T) {
	if !canBuild("aab", "baa") || canBuild("aa", "ab") || !canBuild("", "x") {
		t.Errorf("canBuild works incorrectly")
	}
}
